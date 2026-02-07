import Link from 'next/link';

export const metadata = {
  title: 'Блог'
};

const posts = [
  {
    slug: 'ai-assistants-best-practices',
    title: 'Как проектировать AI‑ассистентов: практики для продакшена',
    excerpt: 'Инструкции, ограничения, контроль качества и измеримость результата.',
    date: '2026-02-01'
  },
  {
    slug: 'telegram-webhook-security',
    title: 'Безопасный Telegram webhook: JWT, секреты и проверка запросов',
    excerpt: 'Разбираем типовые угрозы и базовые меры защиты для микросервисной схемы.',
    date: '2026-01-18'
  },
  {
    slug: 'analytics-engagement-rate',
    title: 'Engagement rate: как считать и улучшать вовлечённость',
    excerpt: 'Метрики, когортный анализ и как не обмануть себя цифрами.',
    date: '2025-12-10'
  }
];

export default function BlogPage() {
  return (
    <div className="mx-auto max-w-7xl container-px py-14">
      <div className="max-w-2xl">
        <h1 className="text-3xl font-semibold tracking-tight">Блог</h1>
        <p className="mt-3 text-slate-600">Статьи об AI, автоматизации и аналитике.</p>
      </div>

      <div className="mt-10 grid grid-cols-1 gap-4 lg:grid-cols-3">
        {posts.map((p) => (
          <Link
            key={p.slug}
            href="#"
            className="rounded-3xl border border-slate-200 bg-white p-7 shadow-sm transition hover:-translate-y-0.5 hover:shadow-md"
          >
            <div className="text-xs font-medium text-slate-500">{p.date}</div>
            <div className="mt-2 text-lg font-semibold tracking-tight">{p.title}</div>
            <div className="mt-2 text-sm text-slate-600">{p.excerpt}</div>
            <div className="mt-6 text-sm font-medium text-brand-700">Читать</div>
          </Link>
        ))}
      </div>
    </div>
  );
}
