'use client';

import { Search, Plus, MoreHorizontal, Circle } from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Switch } from '@/components/ui/switch';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import type { Assistant } from '@/types';

const mockAgents: Assistant[] = [
  {
    id: '1',
    name: 'Support Assistant',
    status: 'active',
    language: 'English',
    timezone: 'UTC',
    crm_sync: true,
    system_prompt: 'You are a helpful support assistant...'
  },
  {
    id: '2',
    name: 'Sales Bot',
    status: 'active',
    language: 'Spanish',
    timezone: 'America/New_York',
    crm_sync: false,
    system_prompt: 'You are a sales assistant...'
  },
  {
    id: '3',
    name: 'FAQ Assistant',
    status: 'training',
    language: 'French',
    timezone: 'Europe/Paris',
    crm_sync: true,
    system_prompt: 'You answer frequently asked questions...'
  },
  {
    id: '4',
    name: 'Technical Support',
    status: 'inactive',
    language: 'German',
    timezone: 'Europe/Berlin',
    crm_sync: false,
    system_prompt: 'Technical support specialist...'
  }
];

const metrics = [
  { label: 'Active Agents', value: '2', change: '+1 from last week' },
  { label: 'Total Conversations', value: '1,234', change: '+12% from last month' },
  { label: 'Customers Engaged', value: '892', change: '+8% from last month' },
  { label: 'Daily Activity', value: '456', change: '+23% from yesterday' }
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

function getStatusBadgeVariant(status: Assistant['status']) {
  switch (status) {
    case 'active':
      return 'default';
    case 'training':
      return 'secondary';
    case 'inactive':
      return 'outline';
    default:
      return 'outline';
  }
}

export default function AgentsPage() {
  return (
    <div className="mx-auto max-w-7xl container-px py-8">
      <div className="flex flex-col gap-8">
        {/* Header */}
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h1 className="text-3xl font-semibold tracking-tight">Agents</h1>
            <p className="mt-2 text-slate-600">Manage your AI assistants and their configurations</p>
          </div>
          <Button asChild>
            <Link href="/agents/new">
              <Plus className="mr-2 h-4 w-4" />
              New Agent
            </Link>
          </Button>
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

        {/* Action Bar */}
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div className="relative flex-1 max-w-sm">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
            <Input
              placeholder="Search agents..."
              className="pl-10"
            />
          </div>
        </div>

        {/* Agents Table */}
        <Card className="p-6">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Status</TableHead>
                <TableHead>Language</TableHead>
                <TableHead>Timezone</TableHead>
                <TableHead>CRM Sync</TableHead>
                <TableHead className="text-right">Actions</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {mockAgents.map((agent) => (
                <TableRow key={agent.id}>
                  <TableCell className="font-medium">{agent.name}</TableCell>
                  <TableCell>
                    <div className="flex items-center gap-2">
                      <Circle className={`h-2 w-2 fill-current ${getStatusColor(agent.status)}`} />
                      <Badge variant={getStatusBadgeVariant(agent.status)}>
                        {agent.status}
                      </Badge>
                    </div>
                  </TableCell>
                  <TableCell>{agent.language}</TableCell>
                  <TableCell>{agent.timezone}</TableCell>
                  <TableCell>
                    <Switch checked={agent.crm_sync} />
                  </TableCell>
                  <TableCell className="text-right">
                    <Button variant="ghost" size="sm">
                      <MoreHorizontal className="h-4 w-4" />
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
