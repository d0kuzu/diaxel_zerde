import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { FeatureCard } from '@/components/feature-card';

export default function HomePage() {
  return (
    <div>
      <section className="relative overflow-hidden">
        <div className="absolute inset-0 -z-10">
          <div className="absolute -top-40 left-1/2 h-[520px] w-[520px] -translate-x-1/2 rounded-full bg-brand-200/60 blur-3xl" />
          <div className="absolute -bottom-40 right-[-120px] h-[520px] w-[520px] rounded-full bg-accent-400/30 blur-3xl" />
          <div className="absolute inset-0 bg-[radial-gradient(40%_60%_at_50%_0%,rgba(59,130,246,0.18),transparent_70%)]" />
        </div>

        <div className="mx-auto max-w-7xl container-px py-16 sm:py-20 lg:py-24">
          <div className="mx-auto max-w-3xl text-center">
            <div className="flex justify-center">
              <Badge>Платформа управления AI‑ассистентами</Badge>
            </div>

            <h1 className="mt-6 text-balance text-4xl font-semibold tracking-tight sm:text-5xl">
              AI‑ассистенты для автоматизации общения и поддержки
            </h1>
            <p className="mt-5 text-pretty text-lg text-slate-600">
              Автопланирование, мультиканальная интеграция и прозрачная аналитика —
              создавайте и развивайте ботов без глубоких знаний ML.
            </p>

            <div className="mt-8 flex flex-col items-center justify-center gap-3 sm:flex-row">
              <Button asChild variant="primary">
                <Link href="/analytics">Попробовать демо</Link>
              </Button>
              <Button asChild variant="secondary">
                <Link href="/pricing">Начать сейчас</Link>
              </Button>
              <Button asChild variant="outline">
                <Link href="/bot-test">Тест ботов</Link>
              </Button>
            </div>

            <div className="mt-10 grid grid-cols-1 gap-4 sm:grid-cols-3">
              <div className="rounded-2xl border border-slate-200 bg-white/70 p-5 shadow-sm backdrop-blur">
                <div className="text-sm font-medium text-slate-900">Интеграции</div>
                <div className="mt-1 text-sm text-slate-600">Telegram, Web‑чат (и расширение далее)</div>
              </div>
              <div className="rounded-2xl border border-slate-200 bg-white/70 p-5 shadow-sm backdrop-blur">
                <div className="text-sm font-medium text-slate-900">Контроль поведения</div>
                <div className="mt-1 text-sm text-slate-600">Инструкции, тональность, сценарии</div>
              </div>
              <div className="rounded-2xl border border-slate-200 bg-white/70 p-5 shadow-sm backdrop-blur">
                <div className="text-sm font-medium text-slate-900">Аналитика</div>
                <div className="mt-1 text-sm text-slate-600">Чаты, пользователи, вовлечённость</div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section className="mx-auto max-w-7xl container-px py-16">
        <div className="flex items-end justify-between gap-6">
          <div>
            <h2 className="text-2xl font-semibold tracking-tight">Функции, которые ускоряют запуск</h2>
            <p className="mt-2 text-slate-600">
              Всё необходимое для управления ассистентами и прозрачной работы с чатами.
            </p>
          </div>
          <div className="hidden sm:block">
            <Button asChild variant="ghost">
              <Link href="/features">Смотреть все</Link>
            </Button>
          </div>
        </div>

        <div className="mt-8 grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
          <FeatureCard
            title="Управление ассистентами"
            description="Создавайте ассистентов, версии инструкций и режимы работы под разные сценарии."
          />
          <FeatureCard
            title="История сообщений"
            description="Контроль диалогов, статусы, поиск и выгрузка. Удобно для QA и улучшений."
          />
          <FeatureCard
            title="Панель аналитики"
            description="Чаты, активные пользователи, engagement rate и фильтры по ассистентам/каналам."
          />
        </div>

        <div className="mt-8 sm:hidden">
          <Button asChild variant="ghost" className="w-full">
            <Link href="/features">Смотреть все</Link>
          </Button>
        </div>
      </section>

      <section className="mx-auto max-w-7xl container-px pb-20">
        <div className="rounded-3xl border border-slate-200 bg-gradient-to-br from-brand-50 via-white to-accent-400/10 p-8 shadow-sm sm:p-10">
          <div className="grid grid-cols-1 gap-8 lg:grid-cols-2 lg:items-center">
            <div>
              <h3 className="text-2xl font-semibold tracking-tight">
                Подключите Telegram webhook и управляйте доступом через JWT
              </h3>
              <p className="mt-3 text-slate-600">
                Архитектура готова к микросервисам: API Gateway, Auth Service, AI Service,
                Database Service. Логи запросов и ошибок — по умолчанию.
              </p>
            </div>
            <div className="flex flex-col gap-3 sm:flex-row sm:justify-end">
              <Button asChild variant="primary">
                <Link href="/contacts">Связаться</Link>
              </Button>
              <Button asChild variant="secondary">
                <Link href="/blog">Читать блог</Link>
              </Button>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}
