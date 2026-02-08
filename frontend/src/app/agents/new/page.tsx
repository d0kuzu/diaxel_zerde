'use client';

import { useState } from 'react';
import { ArrowLeft, Save } from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { PasswordInput } from '@/components/ui/password-input';
import { Textarea } from '@/components/ui/textarea';
import { Switch } from '@/components/ui/switch';
import { AlertCircle } from 'lucide-react';
import { saveAssistantToken } from '@/lib/api';
import { validateTelegramBotToken } from '@/lib/validation';
import type { Assistant } from '@/types';

const newAgent: Omit<Assistant, 'id' | 'bot_token'> = {
  name: '',
  status: 'inactive',
  language: 'English',
  timezone: 'UTC',
  crm_sync: false,
  system_prompt: 'You are a helpful AI assistant.'
};

export default function NewAgentPage() {
  const [agent, setAgent] = useState(newAgent);
  const [botToken, setBotToken] = useState('');
  const [isSaving, setIsSaving] = useState(false);
  const [tokenError, setTokenError] = useState<string | null>(null);
  const [saveMessage, setSaveMessage] = useState<{ type: 'success' | 'error'; message: string } | null>(null);

  const handleSave = async () => {
    // Validate agent name
    if (!agent.name.trim()) {
      setSaveMessage({ type: 'error', message: 'Agent name is required' });
      return;
    }

    // Validate token if provided
    if (botToken.trim()) {
      const validation = validateTelegramBotToken(botToken);
      if (!validation.isValid) {
        setTokenError(validation.message || 'Invalid token format');
        return;
      }
    }

    setTokenError(null);
    setIsSaving(true);

    try {
      // Generate a temporary ID for the new agent
      const tempId = Date.now().toString();
      
      // Save bot token if provided
      if (botToken.trim()) {
        const response = await saveAssistantToken({
          assistant_id: tempId,
          bot_token: botToken.trim()
        });

        if (!response.success) {
          setSaveMessage({ type: 'error', message: 'Failed to save Telegram token' });
          return;
        }
      }

      setSaveMessage({ 
        type: 'success', 
        message: 'Agent created successfully! Redirecting...' 
      });

      // Redirect to agent details page after a short delay
      setTimeout(() => {
        window.location.href = `/agents/${tempId}`;
      }, 1500);

    } catch (error) {
      console.error('Failed to create agent:', error);
      setSaveMessage({ type: 'error', message: 'Failed to create agent. Please try again.' });
    } finally {
      setIsSaving(false);
      setTimeout(() => setSaveMessage(null), 3000);
    }
  };

  return (
    <div className="mx-auto max-w-7xl container-px py-8">
      <div className="flex flex-col gap-8">
        {/* Header */}
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div className="flex items-center gap-4">
            <Button variant="ghost" size="sm" asChild>
              <Link href="/agents">
                <ArrowLeft className="mr-2 h-4 w-4" />
                Back to Agents
              </Link>
            </Button>
            <div>
              <h1 className="text-3xl font-semibold tracking-tight">Create New Agent</h1>
              <p className="mt-2 text-slate-600">Set up a new AI assistant with Telegram integration</p>
            </div>
          </div>
          <Button onClick={handleSave} disabled={isSaving}>
            <Save className="mr-2 h-4 w-4" />
            {isSaving ? 'Creating...' : 'Create Agent'}
          </Button>
        </div>

        {/* Form */}
        <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
          <Card className="p-6">
            <h2 className="text-xl font-semibold text-slate-900 mb-6">Basic Information</h2>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Agent Name *
                </label>
                <Input
                  value={agent.name}
                  onChange={(e) => setAgent({ ...agent, name: e.target.value })}
                  placeholder="Enter agent name..."
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Language
                </label>
                <Input 
                  value={agent.language}
                  onChange={(e) => setAgent({ ...agent, language: e.target.value })}
                  placeholder="English"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Timezone
                </label>
                <Input 
                  value={agent.timezone}
                  onChange={(e) => setAgent({ ...agent, timezone: e.target.value })}
                  placeholder="UTC"
                />
              </div>
            </div>
          </Card>

          <Card className="p-6">
            <h2 className="text-xl font-semibold text-slate-900 mb-6">Telegram Integration</h2>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Telegram Bot Token
                </label>
                <PasswordInput
                  value={botToken}
                  onChange={(e) => {
                    setBotToken(e.target.value);
                    setTokenError(null);
                  }}
                  placeholder="123456789:ABCdefGHIjklmnoPQRstuVWXyz..."
                />
                {tokenError && (
                  <div className="mt-2 flex items-center gap-2 text-sm text-red-600">
                    <AlertCircle className="h-4 w-4" />
                    {tokenError}
                  </div>
                )}
                <p className="mt-2 text-xs text-slate-500">
                  Optional: Enter your Telegram Bot token to enable Telegram integration
                </p>
              </div>
              <div className="flex items-center justify-between">
                <div>
                  <label className="text-sm font-medium text-slate-700">
                    CRM Sync
                  </label>
                  <p className="text-sm text-slate-500">
                    Sync conversations with your CRM system
                  </p>
                </div>
                <Switch
                  checked={agent.crm_sync}
                  onCheckedChange={(checked) => setAgent({ ...agent, crm_sync: checked })}
                />
              </div>
              {saveMessage && (
                <div className={`p-3 rounded-lg text-sm ${
                  saveMessage.type === 'success' 
                    ? 'bg-green-50 text-green-700 border border-green-200' 
                    : 'bg-red-50 text-red-700 border border-red-200'
                }`}>
                  {saveMessage.message}
                </div>
              )}
            </div>
          </Card>
        </div>

        <Card className="p-6">
          <h2 className="text-xl font-semibold text-slate-900 mb-6">System Prompt</h2>
          <Textarea
            value={agent.system_prompt}
            onChange={(e) => setAgent({ ...agent, system_prompt: e.target.value })}
            className="min-h-[200px] resize-none"
            placeholder="Define the AI assistant's behavior, role, and instructions..."
          />
          <p className="mt-2 text-xs text-slate-500">
            Provide clear instructions for how the AI assistant should behave and respond
          </p>
        </Card>
      </div>
    </div>
  );
}
