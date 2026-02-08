'use client';

import { useState, useRef, useEffect } from 'react';
import { ArrowLeft, Send, User, Bot, Phone, Mail, Calendar, Archive, RefreshCw, MessageCircle, Search, Paperclip, Smile, Settings, Info, MoreVertical } from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Card } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Badge } from '@/components/ui/badge';
import { Textarea } from '@/components/ui/textarea';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Separator } from '@/components/ui/separator';
import type { Conversation, Message } from '@/types';

const mockConversation: Conversation = {
  id: '1',
  customer_identity: '+1 (555) 123-4567',
  status: 'active',
  message_count: 24,
  last_activity: '2024-01-15T14:30:00Z',
  assigned_agent: 'Support Assistant'
};

const mockMessages: Message[] = [
  {
    id: '1',
    role: 'user',
    content: 'Hi, I need help with setting up my AI assistant. Can you guide me through the process?',
    timestamp: '2024-01-15T14:00:00Z'
  },
  {
    id: '2',
    role: 'assistant',
    content: 'Hello! I\'d be happy to help you set up your AI assistant. Let me guide you through the process step by step.\n\nFirst, could you tell me what type of assistance you\'re looking to provide? For example:\n- Customer support\n- Sales assistance\n- FAQ automation\n- Technical support\n\nThis will help me recommend the best configuration for your needs.',
    timestamp: '2024-01-15T14:02:00Z'
  },
  {
    id: '3',
    role: 'user',
    content: 'I need customer support automation for my e-commerce store. We get a lot of questions about order status and returns.',
    timestamp: '2024-01-15T14:05:00Z'
  },
  {
    id: '4',
    role: 'assistant',
    content: 'Perfect! Customer support automation for e-commerce is one of our most common use cases. Here\'s what we\'ll set up:\n\n1. **Order Status Integration**: Connect your order management system\n2. **Returns Process**: Automate return requests and status updates\n3. **Product Information**: Access to your product catalog\n4. **Escalation Rules**: When to transfer to human agents\n\nLet\'s start with the basics. What e-commerce platform are you using? (Shopify, WooCommerce, Magento, etc.)',
    timestamp: '2024-01-15T14:07:00Z'
  },
  {
    id: '5',
    role: 'user',
    content: 'We\'re using Shopify. Also, we want to handle questions in both English and Spanish.',
    timestamp: '2024-01-15T14:10:00Z'
  },
  {
    id: '6',
    role: 'assistant',
    content: 'Excellent! Shopify integration is fully supported, and we can definitely handle both English and Spanish.\n\nFor Shopify, we can:\n- ✅ Real-time order status lookup\n- ✅ Product information and inventory\n- ✅ Return initiation and tracking\n- ✅ Customer account information\n\nFor multilingual support:\n- ✅ Automatic language detection\n- ✅ Native responses in both languages\n- ✅ Consistent tone and terminology\n\nWould you like me to start with a basic setup and then we can customize the responses for your specific products and policies?',
    timestamp: '2024-01-15T14:12:00Z'
  }
];

function formatTime(dateString: string) {
  const date = new Date(dateString);
  return date.toLocaleTimeString('en-US', { 
    hour: '2-digit', 
    minute: '2-digit',
    hour12: true 
  });
}

function formatDate(dateString: string) {
  const date = new Date(dateString);
  return date.toLocaleDateString('en-US', { 
    month: 'short', 
    day: 'numeric',
    year: 'numeric'
  });
}

export default function ChatInterfacePage({ params }: { params: Promise<{ id: string }> }) {
  const [messages, setMessages] = useState(mockMessages);
  const [newMessage, setNewMessage] = useState('');
  const [isTyping, setIsTyping] = useState(false);
  const [resolvedParams, setResolvedParams] = useState<{ id: string } | null>(null);

  // Resolve params promise
  useEffect(() => {
    params.then(setResolvedParams);
  }, [params]);
  const [searchQuery, setSearchQuery] = useState('');
  const [showCustomerInfo, setShowCustomerInfo] = useState(true);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const handleSendMessage = () => {
    if (newMessage.trim()) {
      const userMessage: Message = {
        id: Date.now().toString(),
        role: 'user',
        content: newMessage,
        timestamp: new Date().toISOString()
      };
      
      setMessages([...messages, userMessage]);
      setNewMessage('');
      setIsTyping(true);

      // Simulate assistant response
      setTimeout(() => {
        const assistantMessage: Message = {
          id: (Date.now() + 1).toString(),
          role: 'assistant',
          content: 'I understand your message. Let me help you with that...',
          timestamp: new Date().toISOString()
        };
        setMessages(prev => [...prev, assistantMessage]);
        setIsTyping(false);
      }, 1500);
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  function groupMessagesByDate(messages: Message[]) {
    const groups: { [date: string]: Message[] } = {};
    
    messages.forEach(message => {
      const date = formatDate(message.timestamp);
      if (!groups[date]) {
        groups[date] = [];
      }
      groups[date].push(message);
    });
    
    return groups;
  }

  const messageGroups = groupMessagesByDate(messages);

  return (
    <div className="flex h-screen bg-slate-50">
      {/* Left Sidebar - Customer Info & Search */}
      <div className="w-80 bg-white border-r border-slate-200 flex flex-col">
        {/* Search Bar */}
        <div className="p-4 border-b border-slate-200">
          <div className="relative">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-slate-400" />
            <Input
              placeholder="Search conversations..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="pl-10"
            />
          </div>
        </div>

        {/* Customer Information */}
        <div className="flex-1 overflow-y-auto">
          <div className="p-4">
            <div className="flex items-center justify-between mb-4">
              <h3 className="font-semibold text-slate-900">Customer Information</h3>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setShowCustomerInfo(!showCustomerInfo)}
              >
                <Info className="h-4 w-4" />
              </Button>
            </div>
            
            {showCustomerInfo && (
              <div className="space-y-4">
                <div className="flex items-center gap-3">
                  <Avatar className="h-12 w-12">
                    <AvatarFallback>JD</AvatarFallback>
                  </Avatar>
                  <div>
                    <p className="font-medium text-slate-900">John Doe</p>
                    <p className="text-sm text-slate-500">{mockConversation.customer_identity}</p>
                  </div>
                </div>

                <Separator />

                <div className="space-y-3">
                  <div>
                    <p className="text-sm font-medium text-slate-700">Email</p>
                    <p className="text-sm text-slate-600">john.doe@example.com</p>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-slate-700">Customer ID</p>
                    <p className="text-sm text-slate-600">CUST-001234</p>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-slate-700">Status</p>
                    <Badge variant="default" className="mt-1">Premium</Badge>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-slate-700">Total Orders</p>
                    <p className="text-sm text-slate-600">47</p>
                  </div>
                </div>

                <Separator />

                <div className="space-y-3">
                  <div>
                    <p className="text-sm font-medium text-slate-700">Conversation Status</p>
                    <Badge variant="default" className="mt-1">Active</Badge>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-slate-700">Assigned Agent</p>
                    <p className="text-sm text-slate-600">{mockConversation.assigned_agent}</p>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-slate-700">Started</p>
                    <p className="text-sm text-slate-600">{formatDate(mockConversation.last_activity)}</p>
                  </div>
                </div>

                <Separator />

                <div className="space-y-2">
                  <Button variant="outline" className="w-full justify-start">
                    <Phone className="mr-2 h-4 w-4" />
                    Call Customer
                  </Button>
                  <Button variant="outline" className="w-full justify-start">
                    <Settings className="mr-2 h-4 w-4" />
                    Conversation Settings
                  </Button>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Center - Chat Messages */}
      <div className="flex-1 flex flex-col">
        {/* Chat Header */}
        <div className="bg-white border-b border-slate-200 px-6 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <Button variant="ghost" size="sm" asChild>
                <Link href="/conversations">
                  <ArrowLeft className="mr-2 h-4 w-4" />
                  Back
                </Link>
              </Button>
              <Avatar className="h-10 w-10">
                <AvatarFallback>JD</AvatarFallback>
              </Avatar>
              <div>
                <p className="font-semibold text-slate-900">John Doe</p>
                <p className="text-sm text-slate-500">{mockConversation.customer_identity}</p>
              </div>
            </div>
            <div className="flex items-center gap-2">
              <Badge variant="default" className="bg-green-100 text-green-800">
                Active
              </Badge>
              <Button variant="ghost" size="sm">
                <Phone className="h-4 w-4" />
              </Button>
              <Button variant="ghost" size="sm">
                <MoreVertical className="h-4 w-4" />
              </Button>
            </div>
          </div>
        </div>

        {/* Messages Area */}
        <div className="flex-1 overflow-y-auto px-6 py-4">
          <div className="max-w-3xl mx-auto space-y-4">
            {Object.entries(messageGroups).map(([date, msgs]) => (
              <div key={date}>
                <div className="flex items-center justify-center my-4">
                  <div className="bg-slate-100 px-3 py-1 rounded-full">
                    <p className="text-xs font-medium text-slate-600">{date}</p>
                  </div>
                </div>
                {msgs.map((msg) => (
                  <div
                    key={msg.id}
                    className={`flex ${msg.role === 'user' ? 'justify-end' : 'justify-start'} mb-4`}
                  >
                    <div className={`flex gap-3 max-w-lg ${msg.role === 'user' ? 'flex-row-reverse' : 'flex-row'}`}>
                      <Avatar className="h-8 w-8 flex-shrink-0">
                        <AvatarFallback>
                          {msg.role === 'user' ? 'U' : 'A'}
                        </AvatarFallback>
                      </Avatar>
                      <div className={`space-y-1 ${msg.role === 'user' ? 'items-end' : 'items-start'}`}>
                        <div
                          className={`px-4 py-2 rounded-lg ${
                            msg.role === 'user'
                              ? 'bg-brand-600 text-white'
                              : 'bg-slate-100 text-slate-900'
                          }`}
                        >
                          <p className="text-sm whitespace-pre-wrap">{msg.content}</p>
                        </div>
                        <p className="text-xs text-slate-500 px-1">
                          {formatTime(msg.timestamp)}
                        </p>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            ))}
            
            {isTyping && (
              <div className="flex gap-3 justify-start">
                <Avatar className="h-8 w-8 flex-shrink-0">
                  <AvatarFallback>A</AvatarFallback>
                </Avatar>
                <div className="bg-slate-100 rounded-lg px-4 py-3">
                  <div className="flex gap-1">
                    <div className="h-2 w-2 bg-slate-400 rounded-full animate-bounce" />
                    <div className="h-2 w-2 bg-slate-400 rounded-full animate-bounce delay-100" />
                    <div className="h-2 w-2 bg-slate-400 rounded-full animate-bounce delay-200" />
                  </div>
                </div>
              </div>
            )}
            <div ref={messagesEndRef} />
          </div>
        </div>

        {/* Message Input */}
        <div className="bg-white border-t border-slate-200 px-6 py-4">
          <div className="flex items-end gap-3">
            <Button variant="ghost" size="sm" className="p-2">
              <Paperclip className="h-5 w-5 text-slate-400" />
            </Button>
            <div className="flex-1">
              <Textarea
                value={newMessage}
                onChange={(e) => setNewMessage(e.target.value)}
                placeholder="Type your message..."
                className="min-h-[44px] resize-none"
                onKeyDown={handleKeyPress}
              />
            </div>
            <Button variant="ghost" size="sm" className="p-2">
              <Smile className="h-5 w-5 text-slate-400" />
            </Button>
            <Button onClick={handleSendMessage} disabled={!newMessage.trim()}>
              <Send className="h-4 w-4 mr-2" />
              Send
            </Button>
          </div>
        </div>
      </div>

      {/* Right Sidebar - Quick Actions & Templates */}
      <div className="w-80 bg-white border-l border-slate-200 flex flex-col">
        <div className="p-4 border-b border-slate-200">
          <h3 className="font-semibold text-slate-900">Quick Actions</h3>
        </div>
        
        <div className="flex-1 overflow-y-auto p-4">
          <div className="space-y-4">
            {/* Quick Responses */}
            <div>
              <h4 className="text-sm font-medium text-slate-700 mb-3">Quick Responses</h4>
              <div className="space-y-2">
                <Button variant="outline" className="w-full justify-start text-left h-auto p-3">
                  <div>
                    <p className="font-medium text-sm">Thank you for contacting</p>
                    <p className="text-xs text-slate-500">Standard greeting</p>
                  </div>
                </Button>
                <Button variant="outline" className="w-full justify-start text-left h-auto p-3">
                  <div>
                    <p className="font-medium text-sm">I'll check that for you</p>
                    <p className="text-xs text-slate-500">When investigating</p>
                  </div>
                </Button>
                <Button variant="outline" className="w-full justify-start text-left h-auto p-3">
                  <div>
                    <p className="font-medium text-sm">Is there anything else?</p>
                    <p className="text-xs text-slate-500">Closing statement</p>
                  </div>
                </Button>
              </div>
            </div>

            <Separator />

            {/* Order Information */}
            <div>
              <h4 className="text-sm font-medium text-slate-700 mb-3">Recent Orders</h4>
              <div className="space-y-3">
                <Card className="p-3">
                  <div className="flex justify-between items-start mb-2">
                    <div>
                      <p className="font-medium text-sm">ORD-2024-1234</p>
                      <p className="text-xs text-slate-500">Jan 12, 2024</p>
                    </div>
                    <Badge variant="outline" className="text-xs">In Transit</Badge>
                  </div>
                  <p className="text-xs text-slate-600">Wireless Headphones - $89.99</p>
                </Card>
                <Card className="p-3">
                  <div className="flex justify-between items-start mb-2">
                    <div>
                      <p className="font-medium text-sm">ORD-2024-0987</p>
                      <p className="text-xs text-slate-500">Jan 5, 2024</p>
                    </div>
                    <Badge variant="outline" className="text-xs">Delivered</Badge>
                  </div>
                  <p className="text-xs text-slate-600">USB-C Cable - $12.99</p>
                </Card>
              </div>
            </div>

            <Separator />

            {/* Notes */}
            <div>
              <h4 className="text-sm font-medium text-slate-700 mb-3">Conversation Notes</h4>
              <div className="space-y-2">
                <div className="p-3 bg-slate-50 rounded-lg">
                  <p className="text-xs text-slate-600">
                    <strong>Jan 15, 2:28 PM:</strong> Customer asking about AI assistant setup for e-commerce store.
                  </p>
                </div>
                <div className="p-3 bg-slate-50 rounded-lg">
                  <p className="text-xs text-slate-600">
                    <strong>Jan 15, 2:25 PM:</strong> Customer confirmed Shopify platform and bilingual support needed.
                  </p>
                </div>
              </div>
            </div>

            <Separator />

            {/* Actions */}
            <div>
              <h4 className="text-sm font-medium text-slate-700 mb-3">Actions</h4>
              <div className="space-y-2">
                <Button variant="secondary" className="w-full justify-start">
                  Sync to CRM
                </Button>
                <Button variant="secondary" className="w-full justify-start">
                  View History
                </Button>
                <Button variant="secondary" className="w-full justify-start">
                  Transfer to Human
                </Button>
                <Button variant="secondary" className="w-full justify-start">
                  Add Tags
                </Button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
