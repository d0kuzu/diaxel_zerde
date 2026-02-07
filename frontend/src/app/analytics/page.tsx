import { Card } from '@/components/ui/card';
import { AnalyticsDashboard } from './ui/analytics-dashboard';

export const metadata = {
  title: 'Аналитика'
};

export default function AnalyticsPage() {
  return (
    <div className="mx-auto max-w-7xl container-px py-14">
      <div className="flex flex-col gap-2">
        <h1 className="text-3xl font-semibold tracking-tight">Аналитика</h1>
        <p className="text-slate-600">
          Визуализация по чатам, пользователям и engagement rate. Фильтры по ассистенту,
          платформе и дате.
        </p>
      </div>

      <div className="mt-8">
        <Card className="p-0">
          <AnalyticsDashboard />
        </Card>
      </div>
    </div>
  );
}
