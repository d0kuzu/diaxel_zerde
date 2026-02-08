import Link from 'next/link';
import { Button } from '@/components/ui/button';

const links = [
  { href: '/features', label: 'Функции' },
  { href: '/analytics', label: 'Аналитика' },
  { href: '/pricing', label: 'Тарифы' },
  { href: '/blog', label: 'Блог' },
  { href: '/contacts', label: 'Контакты' }
];

export function Navbar() {
  return (
    <header className="sticky top-0 z-50 border-b border-slate-200 bg-white/70 backdrop-blur">
      <div className="mx-auto flex max-w-7xl items-center justify-between gap-4 container-px py-4">
        <Link href="/" className="flex items-center gap-2">
          <img 
            src="/large(1).ico" 
            alt="SD Nexus" 
            className="h-9 w-9 rounded-xl"
          />
          <span className="text-sm font-semibold tracking-tight">SD Nexus</span>
        </Link>

        <nav className="hidden items-center gap-6 md:flex">
          {links.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              className="text-sm text-slate-600 transition-colors hover:text-slate-900"
            >
              {l.label}
            </Link>
          ))}
        </nav>

        <div className="flex items-center gap-2">
          <Button asChild variant="ghost" className="hidden sm:inline-flex">
            <Link href="/analytics">Демо</Link>
          </Button>
          <Button asChild variant="ghost" className="hidden sm:inline-flex">
            <Link href="/login">Войти</Link>
          </Button>
          <Button asChild variant="primary">
            <Link href="/register">Регистрация</Link>
          </Button>
        </div>
      </div>
    </header>
  );
}
