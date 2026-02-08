export interface Assistant {
  id: string;
  name: string;
  status: 'active' | 'inactive' | 'training';
  language: string;
  timezone: string;
  crm_sync: boolean;
  system_prompt: string;
  bot_token?: string;
}

export interface Conversation {
  id: string;
  customer_identity: string;
  status: 'active' | 'closed';
  message_count: number;
  last_activity: string;
  assigned_agent?: string;
}

export interface Message {
  id: string;
  role: 'user' | 'assistant';
  content: string;
  timestamp: string;
}
