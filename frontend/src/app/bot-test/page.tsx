'use client';

import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Card } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';

export default function BotTestPage() {
  const [assistantId, setAssistantId] = useState('test-assistant-123');
  const [botToken, setBotToken] = useState('1234567890:ABCdefGHIjklMNOpqrsTUVwxyz');
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState<any>(null);
  const [error, setError] = useState<string>('');

  const API_BASE = 'http://localhost:8080'; // assuming auth-service runs on 8080

  const createAssistant = async () => {
    setLoading(true);
    setError('');
    setResult(null);

    try {
      const response = await fetch(`${API_BASE}/assistant`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          assistant_id: assistantId,
          bot_token: botToken,
        }),
      });

      const data = await response.json();
      setResult(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
    } finally {
      setLoading(false);
    }
  };

  const getBotToken = async () => {
    setLoading(true);
    setError('');
    setResult(null);

    try {
      const response = await fetch(`${API_BASE}/assistant/${assistantId}/bot-token`);
      const data = await response.json();
      setResult(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
    } finally {
      setLoading(false);
    }
  };

  const getTestInfo = async () => {
    setLoading(true);
    setError('');
    setResult(null);

    try {
      const response = await fetch(`${API_BASE}/test/bot-registration`);
      const data = await response.json();
      setResult(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Unknown error');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto p-6 max-w-4xl">
      <div className="mb-8">
        <h1 className="text-3xl font-bold mb-2">Тестирование регистрации Telegram ботов</h1>
        <p className="text-gray-600">Тестовая страница для проверки API эндпоинтов управления ботами</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <Card>
          <div className="mb-4">
            <h3 className="text-lg font-semibold mb-1">Создание/обновление ассистента</h3>
            <p className="text-sm text-gray-600">POST /assistant</p>
          </div>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium mb-2">Assistant ID</label>
              <Input
                value={assistantId}
                onChange={(e) => setAssistantId(e.target.value)}
                placeholder="test-assistant-123"
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-2">Bot Token</label>
              <Input
                value={botToken}
                onChange={(e) => setBotToken(e.target.value)}
                placeholder="1234567890:ABCdefGHIjklMNOpqrsTUVwxyz"
              />
            </div>
            <Button 
              onClick={createAssistant} 
              disabled={loading}
              className="w-full"
            >
              {loading ? 'Загрузка...' : 'Создать/обновить ассистента'}
            </Button>
          </div>
        </Card>

        <Card>
          <div className="mb-4">
            <h3 className="text-lg font-semibold mb-1">Получение токена ассистента</h3>
            <p className="text-sm text-gray-600">GET /assistant/:assistant_id/bot-token</p>
          </div>
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium mb-2">Assistant ID</label>
              <Input
                value={assistantId}
                onChange={(e) => setAssistantId(e.target.value)}
                placeholder="test-assistant-123"
              />
            </div>
            <Button 
              onClick={getBotToken} 
              disabled={loading}
              variant="outline"
              className="w-full"
            >
              {loading ? 'Загрузка...' : 'Получить bot token'}
            </Button>
          </div>
        </Card>
      </div>

      <div className="mb-6">
        <Button onClick={getTestInfo} disabled={loading} variant="secondary">
          {loading ? 'Загрузка...' : 'Получить информацию об эндпоинтах'}
        </Button>
      </div>

      {error && (
        <Card className="mb-6 border-red-200">
          <div className="pt-6">
            <Badge variant="secondary" className="mb-2 bg-red-100 text-red-800">Ошибка</Badge>
            <pre className="text-red-600 whitespace-pre-wrap">{error}</pre>
          </div>
        </Card>
      )}

      {result && (
        <Card>
          <div className="mb-4">
            <h3 className="text-lg font-semibold">Результат</h3>
          </div>
          <div>
            <pre className="bg-gray-100 p-4 rounded-md overflow-auto text-sm">
              {JSON.stringify(result, null, 2)}
            </pre>
          </div>
        </Card>
      )}

      <div className="mt-8">
        <Card>
          <div className="mb-4">
            <h3 className="text-lg font-semibold">Информация об API</h3>
          </div>
          <div className="space-y-2 text-sm">
            <p><strong>Base URL:</strong> {API_BASE}</p>
            <p><strong>Эндпоинты:</strong></p>
            <ul className="list-disc list-inside ml-4 space-y-1">
              <li>POST /assistant - Создание/обновление ассистента с bot_token</li>
              <li>GET /assistant/:assistant_id/bot-token - Получение bot_token по ID</li>
              <li>GET /test/bot-registration - Тестовая информация</li>
            </ul>
            <p className="text-gray-600 mt-4">
              Эти эндпоинты работают без авторизации для тестирования.
            </p>
          </div>
        </Card>
      </div>
    </div>
  );
}
