'use client';

import { useState, useEffect } from 'react';
import { ArrowLeft, Save, Settings, Database, BookOpen, AlertCircle } from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { PasswordInput } from '@/components/ui/password-input';
import { Textarea } from '@/components/ui/textarea';
import { Badge } from '@/components/ui/badge';
import { Switch } from '@/components/ui/switch';
import type { Assistant } from '@/types';
import { saveAssistantToken, getBotToken } from '@/lib/api';
import { validateTelegramBotToken } from '@/lib/validation';

const mockAgent: Assistant = {
  id: '1',
  name: 'Support Assistant',
  status: 'active',
  language: 'English',
  timezone: 'UTC',
  crm_sync: true,
  system_prompt: `You are a helpful support assistant for SD Nexus. Your role is to:

1. Provide clear and accurate information about our AI assistant platform
2. Help users with technical issues and questions
3. Guide users through setup and configuration processes
4. Maintain a professional and friendly tone
5. Escalate complex issues to human support when necessary

Key guidelines:
- Always be patient and understanding
- Provide step-by-step instructions when helping with technical tasks
- Ask clarifying questions if the user's request is unclear
- Offer additional help and resources when appropriate`
};

const tabs = [
  { id: 'prompt', label: 'System Prompt', icon: BookOpen },
  { id: 'settings', label: 'Settings', icon: Settings },
  { id: 'tools', label: 'Tools', icon: Database },
  { id: 'knowledge', label: 'Knowledge Base', icon: BookOpen }
];

function getStatusColor(status: Assistant['status']) {
  switch (status) {
    case 'active':
      return 'bg-green-500';
    case 'training':
      return 'bg-yellow-500';
    case 'inactive':
      return 'bg-slate-400';
    default:
      return 'bg-slate-400';
  }
}

export default function AgentDetailsPage({ params }: { params: Promise<{ id: string }> }) {
  const [activeTab, setActiveTab] = useState('prompt');
  const [agent, setAgent] = useState(mockAgent);
  const [botToken, setBotToken] = useState('');
  const [isLoadingToken, setIsLoadingToken] = useState(false);
  const [isSavingToken, setIsSavingToken] = useState(false);
  const [tokenError, setTokenError] = useState<string | null>(null);
  const [saveMessage, setSaveMessage] = useState<{ type: 'success' | 'error'; message: string } | null>(null);
  const [resolvedParams, setResolvedParams] = useState<{ id: string } | null>(null);

  // Resolve params promise
  useEffect(() => {
    params.then(setResolvedParams);
  }, [params]);

  // Load existing bot token on component mount
  useEffect(() => {
    if (!resolvedParams) return;
    
    const loadBotToken = async () => {
      setIsLoadingToken(true);
      try {
        const response = await getBotToken(resolvedParams.id);
        setBotToken(response.bot_token);
      } catch (error) {
        console.error('Failed to load bot token:', error);
      } finally {
        setIsLoadingToken(false);
      }
    };

    loadBotToken();
  }, [resolvedParams]);

  const handleSaveToken = async () => {
    if (!resolvedParams) return;

    // Validate token format
    const validation = validateTelegramBotToken(botToken);
    if (!validation.isValid) {
      setTokenError(validation.message || 'Invalid token format');
      return;
    }

    setTokenError(null);
    setIsSavingToken(true);

    try {
      const response = await saveAssistantToken({
        assistant_id: resolvedParams.id,
        bot_token: botToken.trim()
      });

      if (response.success) {
        setSaveMessage({ type: 'success', message: 'Telegram Bot Token saved successfully!' });
      } else {
        setSaveMessage({ type: 'error', message: 'Failed to save token. Please try again.' });
      }
    } catch (error) {
      console.error('Failed to save bot token:', error);
      setSaveMessage({ type: 'error', message: 'Failed to save token. Please try again.' });
    } finally {
      setIsSavingToken(false);
      // Clear message after 3 seconds
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
              <h1 className="text-3xl font-semibold tracking-tight">{agent.name}</h1>
              <div className="mt-2 flex items-center gap-2">
                <div className={`h-2 w-2 rounded-full bg-current ${getStatusColor(agent.status)}`} />
                <Badge variant={agent.status === 'active' ? 'default' : 'secondary'}>
                  {agent.status}
                </Badge>
              </div>
            </div>
          </div>
          <Button>
            <Save className="mr-2 h-4 w-4" />
            Save Changes
          </Button>
        </div>

        {/* Navigation Tabs */}
        <div className="border-b border-slate-200">
          <nav className="-mb-px flex space-x-8">
            {tabs.map((tab) => {
              const Icon = tab.icon;
              return (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`flex items-center gap-2 border-b-2 py-4 text-sm font-medium transition-colors ${
                    activeTab === tab.id
                      ? 'border-brand-500 text-brand-600'
                      : 'border-transparent text-slate-600 hover:text-slate-900 hover:border-slate-300'
                  }`}
                >
                  <Icon className="h-4 w-4" />
                  {tab.label}
                </button>
              );
            })}
          </nav>
        </div>

        {/* Tab Content */}
        {activeTab === 'prompt' && (
          <div className="grid grid-cols-1 gap-6 lg:grid-cols-3">
            <div className="lg:col-span-2">
              <Card className="p-6">
                <div className="mb-6">
                  <h2 className="text-xl font-semibold text-slate-900">System Prompt</h2>
                  <p className="mt-2 text-sm text-slate-600">
                    Define the AI assistant's behavior, role, and instructions
                  </p>
                </div>
                <Textarea
                  value={agent.system_prompt}
                  onChange={(e) => setAgent({ ...agent, system_prompt: e.target.value })}
                  className="min-h-[400px] resize-none"
                  placeholder="Enter the system prompt..."
                />
              </Card>
            </div>
            <div className="space-y-6">
              <Card className="p-6">
                <h3 className="text-lg font-semibold text-slate-900 mb-4">Quick Actions</h3>
                <div className="space-y-3">
                  <Button variant="secondary" className="w-full justify-start">
                    Reset to Default
                  </Button>
                  <Button variant="secondary" className="w-full justify-start">
                    Test Prompt
                  </Button>
                  <Button variant="secondary" className="w-full justify-start">
                    View History
                  </Button>
                </div>
              </Card>
              <Card className="p-6">
                <h3 className="text-lg font-semibold text-slate-900 mb-4">Prompt Guidelines</h3>
                <ul className="space-y-2 text-sm text-slate-600">
                  <li>• Be specific about the assistant's role</li>
                  <li>• Include clear behavioral guidelines</li>
                  <li>• Define response format expectations</li>
                  <li>• Set escalation procedures</li>
                  <li>• Specify tone and style requirements</li>
                </ul>
              </Card>
            </div>
          </div>
        )}

        {activeTab === 'settings' && (
          <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
            <Card className="p-6">
              <h2 className="text-xl font-semibold text-slate-900 mb-6">Basic Settings</h2>
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-2">
                    Agent Name
                  </label>
                  <Input
                    value={agent.name}
                    onChange={(e) => setAgent({ ...agent, name: e.target.value })}
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-2">
                    Language
                  </label>
                  <Input value={agent.language} />
                </div>
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-2">
                    Timezone
                  </label>
                  <Input value={agent.timezone} />
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
                    disabled={isLoadingToken}
                  />
                  {tokenError && (
                    <div className="mt-2 flex items-center gap-2 text-sm text-red-600">
                      <AlertCircle className="h-4 w-4" />
                      {tokenError}
                    </div>
                  )}
                  <p className="mt-2 text-xs text-slate-500">
                    Enter your Telegram Bot token to enable Telegram integration
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
                <Button 
                  onClick={handleSaveToken} 
                  disabled={isSavingToken || isLoadingToken}
                  className="w-full"
                >
                  {isSavingToken ? 'Saving...' : 'Save Telegram Token'}
                </Button>
              </div>
            </Card>
          </div>
        )}

        {activeTab === 'tools' && (
          <Card className="p-6">
            <h2 className="text-xl font-semibold text-slate-900 mb-6">Available Tools</h2>
            <div className="text-center py-12 text-slate-500">
              <Database className="mx-auto h-12 w-12 text-slate-400 mb-4" />
              <p>Tools configuration coming soon</p>
            </div>
          </Card>
        )}

        {activeTab === 'knowledge' && (
          <Card className="p-6">
            <h2 className="text-xl font-semibold text-slate-900 mb-6">Knowledge Base</h2>
            <div className="text-center py-12 text-slate-500">
              <BookOpen className="mx-auto h-12 w-12 text-slate-400 mb-4" />
              <p>Knowledge base management coming soon</p>
            </div>
          </Card>
        )}
      </div>
    </div>
  );
}
