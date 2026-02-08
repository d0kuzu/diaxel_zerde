'use client';

import { Search, MessageCircle, User, Clock, Filter } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import type { Conversation } from '@/types';

const mockConversations: Conversation[] = [
  {
    id: '1',
    customer_identity: '+1 (555) 123-4567',
    status: 'active',
    message_count: 24,
    last_activity: '2024-01-15T14:30:00Z',
    assigned_agent: 'Support Assistant'
  },
  {
    id: '2',
    customer_identity: 'user@example.com',
    status: 'active',
    message_count: 15,
    last_activity: '2024-01-15T14:15:00Z',
    assigned_agent: 'Sales Bot'
  },
  {
    id: '3',
    customer_identity: '+44 20 7946 0958',
    status: 'closed',
    message_count: 8,
    last_activity: '2024-01-15T12:45:00Z',
    assigned_agent: 'Support Assistant'
  },
  {
    id: '4',
    customer_identity: 'john.doe@company.com',
    status: 'active',
    message_count: 32,
    last_activity: '2024-01-15T14:45:00Z',
    assigned_agent: 'FAQ Assistant'
  },
  {
    id: '5',
    customer_identity: '+1 (555) 987-6543',
    status: 'closed',
    message_count: 12,
    last_activity: '2024-01-15T11:30:00Z',
    assigned_agent: 'Technical Support'
  }
];

const metrics = [
  { label: 'Active Chats', value: '3', change: '+1 from last hour' },
  { label: 'Closed Today', value: '47', change: '+15% from yesterday' },
  { label: 'Avg Response Time', value: '2.3m', change: '-30s from last week' },
  { label: 'Total Messages', value: '1,892', change: '+8% from yesterday' }
];

function getStatusColor(status: Conversation['status']) {
  return status === 'active' ? 'bg-green-500' : 'bg-slate-400';
}

function getStatusBadgeVariant(status: Conversation['status']) {
  return status === 'active' ? 'default' : 'outline';
}

function formatTimeAgo(dateString: string) {
  const date = new Date(dateString);
  const now = new Date();
  const diffInMinutes = Math.floor((now.getTime() - date.getTime()) / (1000 * 60));
  
  if (diffInMinutes < 1) return 'Just now';
  if (diffInMinutes < 60) return `${diffInMinutes}m ago`;
  if (diffInMinutes < 1440) return `${Math.floor(diffInMinutes / 60)}h ago`;
  return `${Math.floor(diffInMinutes / 1440)}d ago`;
}

export default function ConversationsPage() {
  return (
    <div className="mx-auto max-w-7xl container-px py-8">
      <div className="flex flex-col gap-8">
        {/* Header */}
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h1 className="text-3xl font-semibold tracking-tight">Conversations</h1>
            <p className="mt-2 text-slate-600">Monitor and manage customer conversations</p>
          </div>
          <div className="flex gap-2">
            <Button variant="secondary">
              <Filter className="mr-2 h-4 w-4" />
              Filter
            </Button>
          </div>
        </div>

        {/* Metrics Cards */}
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {metrics.map((metric) => (
            <Card key={metric.label} className="p-6">
              <div className="flex flex-col gap-1">
                <p className="text-sm font-medium text-slate-600">{metric.label}</p>
                <p className="text-2xl font-semibold text-slate-900">{metric.value}</p>
                <p className="text-xs text-slate-500">{metric.change}</p>
              </div>
            </Card>
          ))}
        </div>

        {/* Search and Filter Bar */}
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div className="relative flex-1 max-w-sm">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
            <Input
              placeholder="Search conversations..."
              className="pl-10"
            />
          </div>
          <div className="flex gap-2">
            <Badge variant="outline" className="cursor-pointer hover:bg-slate-50">
              All
            </Badge>
            <Badge variant="default" className="cursor-pointer">
              Active
            </Badge>
            <Badge variant="outline" className="cursor-pointer hover:bg-slate-50">
              Closed
            </Badge>
          </div>
        </div>

        {/* Conversations Table */}
        <Card className="p-6">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Customer</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Agent</TableHead>
                <TableHead>Messages</TableHead>
                <TableHead>Last Activity</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {mockConversations.map((conversation) => (
                <TableRow key={conversation.id} className="cursor-pointer hover:bg-slate-50">
                  <TableCell>
                    <div className="flex items-center gap-3">
                      <div className="flex h-8 w-8 items-center justify-center rounded-full bg-brand-100">
                        <User className="h-4 w-4 text-brand-600" />
                      </div>
                      <div>
                        <p className="font-medium text-slate-900">
                          {conversation.customer_identity}
                        </p>
                        <p className="text-sm text-slate-500">ID: {conversation.id}</p>
                      </div>
                    </div>
                  </TableCell>
                  <TableCell>
                    <div className="flex items-center gap-2">
                      <div className={`h-2 w-2 rounded-full ${getStatusColor(conversation.status)}`} />
                      <Badge variant={getStatusBadgeVariant(conversation.status)}>
                        {conversation.status}
                      </Badge>
                    </div>
                  </TableCell>
                  <TableCell>
                    <div className="flex items-center gap-2">
                      <MessageCircle className="h-4 w-4 text-slate-400" />
                      <span className="text-sm text-slate-900">
                        {conversation.assigned_agent}
                      </span>
                    </div>
                  </TableCell>
                  <TableCell>
                    <div className="flex items-center gap-2">
                      <MessageCircle className="h-4 w-4 text-slate-400" />
                      <span className="text-sm text-slate-900">
                        {conversation.message_count}
                      </span>
                    </div>
                  </TableCell>
                  <TableCell>
                    <div className="flex items-center gap-2">
                      <Clock className="h-4 w-4 text-slate-400" />
                      <span className="text-sm text-slate-900">
                        {formatTimeAgo(conversation.last_activity)}
                      </span>
                    </div>
                  </TableCell>
                  <TableCell className="text-right">
                    <Button variant="ghost" size="sm" asChild>
                      <a href={`/conversations/${conversation.id}`}>View</a>
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </Card>
      </div>
    </div>
  );
}
