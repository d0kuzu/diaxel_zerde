import Link from 'next/link';
import { Button } from '@/components/ui/button';

export const metadata = {
  title: 'Тарифы'
};

const plans = [
  {
    name: 'Starter',
    price: '0₽',
    description: 'Для первых экспериментов и демо.',
    items: ['1 ассистент', 'Telegram webhook', 'Базовая аналитика'],
    cta: 'Начать бесплатно',
    variant: 'secondary' as const
  },
  {
    name: 'Pro',
    price: 'от 4 900₽/мес',
    description: 'Для команд и продуктовых запусков.',
    items: ['До 10 ассистентов', 'Фильтры и сегменты', 'Экспорт данных'],
    cta: 'Подключить',
    variant: 'primary' as const
  },
  {
    name: 'Enterprise',
    price: 'Индивидуально',
    description: 'Для больших нагрузок и требований безопасности.',
    items: ['SLA', 'On‑prem / VPC', 'RBAC и аудит'],
    cta: 'Запросить доступ',
    variant: 'secondary' as const
  }
];

export default function PricingPage() {
  return (
    <div className="mx-auto max-w-7xl container-px py-14">
      <div className="max-w-2xl">
        <h1 className="text-3xl font-semibold tracking-tight">Тарифы</h1>
        <p className="mt-3 text-slate-600">
          Прозрачные планы для разных этапов роста. Начните с демо и переходите на Pro по
          мере масштабирования.
        </p>
      </div>

      <div className="mt-10 grid grid-cols-1 gap-4 lg:grid-cols-3">
        {plans.map((p) => (
          <div
            key={p.name}
            className="rounded-3xl border border-slate-200 bg-white p-7 shadow-sm transition hover:-translate-y-0.5 hover:shadow-md"
          >
            <div className="text-sm font-semibold">{p.name}</div>
            <div className="mt-3 text-3xl font-semibold tracking-tight">{p.price}</div>
            <div className="mt-2 text-sm text-slate-600">{p.description}</div>
            <div className="mt-6 flex flex-col gap-2 text-sm text-slate-700">
              {p.items.map((it) => (
                <div key={it} className="flex items-center gap-2">
                  <span className="h-1.5 w-1.5 rounded-full bg-brand-600" />
                  <span>{it}</span>
                </div>
              ))}
            </div>
            <div className="mt-7">
              <Button asChild variant={p.variant} className="w-full">
                <Link href="/contacts">{p.cta}</Link>
              </Button>
            </div>
          </div>
        ))}
      </div>

      <div className="mt-10 rounded-3xl border border-slate-200 bg-brand-50 p-8">
        <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <div className="text-sm font-semibold">Нужна интеграция под вашу инфраструктуру?</div>
            <div className="mt-1 text-sm text-slate-600">
              Поможем с webhook, JWT, настройкой микросервисов и аналитикой.
            </div>
          </div>
          <Button asChild variant="primary">
            <Link href="/contacts">Связаться</Link>
          </Button>
        </div>
      </div>
    </div>
  );
}
