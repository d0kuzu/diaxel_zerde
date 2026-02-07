import Link from 'next/link';
import { FeatureCard } from '@/components/feature-card';
import { Button } from '@/components/ui/button';

export const metadata = {
  title: 'Функции'
};

export default function FeaturesPage() {
  return (
    <div className="mx-auto max-w-7xl container-px py-14">
      <div className="max-w-2xl">
        <h1 className="text-3xl font-semibold tracking-tight">Функции платформы</h1>
        <p className="mt-3 text-slate-600">
          Управляйте ассистентами, чатами и инструкциями, а затем измеряйте эффективность
          через понятную аналитику.
        </p>
      </div>

      <div className="mt-8 grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <FeatureCard
          title="Управление AI‑ассистентами"
          description="Создание ассистентов, окружения (dev/prod), версии конфигураций и быстрый откат."
        />
        <FeatureCard
          title="Инструкции и поведение"
          description="Гайдлайны, тональность, ограничения, system prompts и контент‑политики."
        />
        <FeatureCard
          title="Чаты и история"
          description="Единая история сообщений, статусы, поиск, выгрузка для контроля качества."
        />
        <FeatureCard
          title="Аналитика"
          description="Активные пользователи, чаты, вовлечённость, с фильтрами по ассистентам и каналам."
        />
        <FeatureCard
          title="Безопасность"
          description="JWT‑доступ, разграничение ролей, безопасные webhooks и логирование запросов."
        />
        <FeatureCard
          title="Интеграции"
          description="Telegram и Web‑чат сейчас, расширение на другие платформы по мере роста."
        />
      </div>

      <div className="mt-10 rounded-3xl border border-slate-200 bg-gradient-to-br from-brand-50 via-white to-accent-400/10 p-8 shadow-sm">
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <div className="text-sm font-semibold">Готово к запуску за 1 день</div>
            <div className="mt-1 text-sm text-slate-600">
              Посмотрите демо аналитики и структуру будущей панели.
            </div>
          </div>
          <Button asChild variant="primary">
            <Link href="/analytics">Открыть демо</Link>
          </Button>
        </div>
      </div>
    </div>
  );
}
