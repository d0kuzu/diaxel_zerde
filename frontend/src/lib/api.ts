export type AnalyticsFilters = {
  assistantId?: string;
  platform?: 'telegram' | 'web';
  from?: string;
  to?: string;
};

export type AnalyticsSummary = {
  chats: number;
  activeUsers: number;
  engagementRate: number;
};

export type AnalyticsTimeseriesPoint = {
  date: string;
  chats: number;
  activeUsers: number;
  engagementRate: number;
};

export interface SaveAssistantTokenRequest {
  assistant_id: string;
  bot_token: string;
}

export interface SaveAssistantTokenResponse {
  success: boolean;
}

export interface GetBotTokenResponse {
  bot_token: string;
}

function buildQuery(filters: AnalyticsFilters) {
  const params = new URLSearchParams();
  if (filters.assistantId) params.set('assistantId', filters.assistantId);
  if (filters.platform) params.set('platform', filters.platform);
  if (filters.from) params.set('from', filters.from);
  if (filters.to) params.set('to', filters.to);
  const qs = params.toString();
  return qs ? `?${qs}` : '';
}

export async function fetchAnalyticsSummary(filters: AnalyticsFilters): Promise<AnalyticsSummary> {
  const base = process.env.NEXT_PUBLIC_API_BASE_URL;
  if (!base) {
    return { chats: 128, activeUsers: 56, engagementRate: 0.42 };
  }

  const res = await fetch(`${base}/api/analytics/summary${buildQuery(filters)}`, {
    cache: 'no-store'
  });

  if (!res.ok) {
    throw new Error(`Analytics summary request failed: ${res.status}`);
  }

  return res.json() as Promise<AnalyticsSummary>;
}

export async function fetchAnalyticsTimeseries(
  filters: AnalyticsFilters
): Promise<AnalyticsTimeseriesPoint[]> {
  const base = process.env.NEXT_PUBLIC_API_BASE_URL;
  if (!base) {
    const today = new Date();
    const data: AnalyticsTimeseriesPoint[] = [];
    for (let i = 13; i >= 0; i--) {
      const d = new Date(today);
      d.setDate(today.getDate() - i);
      const date = d.toISOString().slice(0, 10);
      const chats = 6 + Math.round(Math.random() * 10);
      const activeUsers = 3 + Math.round(Math.random() * 7);
      const engagementRate = Math.max(
        0.18,
        Math.min(0.78, 0.25 + Math.random() * 0.4)
      );
      data.push({ date, chats, activeUsers, engagementRate: Number(engagementRate.toFixed(2)) });
    }
    return data;
  }

  const res = await fetch(`${base}/api/analytics/timeseries${buildQuery(filters)}`, {
    cache: 'no-store'
  });

  if (!res.ok) {
    throw new Error(`Analytics timeseries request failed: ${res.status}`);
  }

  return res.json() as Promise<AnalyticsTimeseriesPoint[]>;
}

export async function saveAssistantToken(
  data: SaveAssistantTokenRequest
): Promise<SaveAssistantTokenResponse> {
  const base = process.env.NEXT_PUBLIC_API_BASE_URL;
  if (!base) {
    // Mock response for development
    return { success: true };
  }

  const res = await fetch(`${base}/api/assistants/token`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
    cache: 'no-store'
  });

  if (!res.ok) {
    throw new Error(`Save assistant token request failed: ${res.status}`);
  }

  return res.json() as Promise<SaveAssistantTokenResponse>;
}

export async function getBotToken(
  assistantId: string
): Promise<GetBotTokenResponse> {
  const base = process.env.NEXT_PUBLIC_API_BASE_URL;
  if (!base) {
    // Mock response for development
    return { bot_token: '123456789:ABCdefGHIjklmnoPQRstuVWXyz' };
  }

  const res = await fetch(`${base}/api/assistants/${assistantId}/token`, {
    cache: 'no-store'
  });

  if (!res.ok) {
    throw new Error(`Get bot token request failed: ${res.status}`);
  }

  return res.json() as Promise<GetBotTokenResponse>;
}
