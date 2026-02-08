'use client';

import { useEffect, useMemo, useState } from 'react';
import { fetchAnalyticsSummary, fetchAnalyticsTimeseries, type AnalyticsFilters } from '@/lib/api';
import { Select } from '@/components/ui/select';
import { Input } from '@/components/ui/input';
import { AnalyticsCharts } from '@/components/charts/analytics-charts';

const assistants = [
  { id: 'assistant-1', name: 'Support Assistant' },
  { id: 'assistant-2', name: 'Sales Assistant' },
  { id: 'assistant-3', name: 'Edu Assistant' }
];

export function AnalyticsDashboard() {
  const [assistantId, setAssistantId] = useState<string>('');
  const [platform, setPlatform] = useState<string>('');
  const [from, setFrom] = useState<string>('');
  const [to, setTo] = useState<string>('');

  const [summary, setSummary] = useState<{ chats: number; activeUsers: number; engagementRate: number } | null>(
    null
  );
  const [series, setSeries] = useState<Array<{ date: string; chats: number; activeUsers: number; engagementRate: number }>>(
    []
  );
  const [error, setError] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(true);

  const filters = useMemo((): AnalyticsFilters => {
    const f: AnalyticsFilters = {};
    if (assistantId) f.assistantId = assistantId;
    if (platform === 'telegram' || platform === 'web') f.platform = platform;
    if (from) f.from = from;
    if (to) f.to = to;
    return f;
  }, [assistantId, platform, from, to]);

  useEffect(() => {
    let alive = true;
    setLoading(true);
    setError('');

    Promise.all([fetchAnalyticsSummary(filters), fetchAnalyticsTimeseries(filters)])
      .then(([s, t]) => {
        if (!alive) return;
        setSummary(s);
        setSeries(t);
      })
      .catch((e: unknown) => {
        if (!alive) return;
        setError(e instanceof Error ? e.message : 'Unknown error');
      })
      .finally(() => {
        if (!alive) return;
        setLoading(false);
      });

    return () => {
      alive = false;
    };
  }, [filters]);

  return (
    <div className="p-6">
      <div className="grid grid-cols-1 gap-3 md:grid-cols-4">
        <div>
          <div className="mb-2 text-xs font-medium text-slate-600">Ассистент</div>
          <Select value={assistantId} onChange={(e) => setAssistantId(e.target.value)}>
            <option value="">Все</option>
            {assistants.map((a) => (
              <option key={a.id} value={a.id}>
                {a.name}
              </option>
            ))}
          </Select>
        </div>

        <div>
          <div className="mb-2 text-xs font-medium text-slate-600">Платформа</div>
          <Select value={platform} onChange={(e) => setPlatform(e.target.value)}>
            <option value="">Все</option>
            <option value="telegram">Telegram</option>
            <option value="web">Web‑чат</option>
          </Select>
        </div>

        <div>
          <div className="mb-2 text-xs font-medium text-slate-600">Дата с</div>
          <Input type="date" value={from} onChange={(e) => setFrom(e.target.value)} />
        </div>

        <div>
          <div className="mb-2 text-xs font-medium text-slate-600">Дата по</div>
          <Input type="date" value={to} onChange={(e) => setTo(e.target.value)} />
        </div>
      </div>

      {error ? (
        <div className="mt-5 rounded-2xl border border-red-200 bg-red-50 p-4 text-sm text-red-800">
          {error}
        </div>
      ) : null}

      <div className="mt-6 grid grid-cols-1 gap-4 sm:grid-cols-3">
        <div className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm">
          <div className="text-xs font-medium text-slate-600">Чаты</div>
          <div className="mt-2 text-2xl font-semibold">
            {loading ? '—' : summary?.chats ?? '—'}
          </div>
        </div>
        <div className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm">
          <div className="text-xs font-medium text-slate-600">Активные пользователи</div>
          <div className="mt-2 text-2xl font-semibold">
            {loading ? '—' : summary?.activeUsers ?? '—'}
          </div>
        </div>
        <div className="rounded-2xl border border-slate-200 bg-white p-5 shadow-sm">
          <div className="text-xs font-medium text-slate-600">Engagement rate</div>
          <div className="mt-2 text-2xl font-semibold">
            {loading ? '—' : `${Math.round((summary?.engagementRate ?? 0) * 100)}%`}
          </div>
        </div>
      </div>

      <div className="mt-6">
        <AnalyticsCharts data={series} />
      </div>

      <div className="mt-6 text-xs text-slate-500">
        Если выставить `NEXT_PUBLIC_API_BASE_URL`, дашборд будет запрашивать реальные данные
        с `/api/analytics/*`.
      </div>
    </div>
  );
}
