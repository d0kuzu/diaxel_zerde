import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';
import { Navbar } from '@/components/navbar';
import { Footer } from '@/components/footer';

import type { ReactNode } from 'react';

const inter = Inter({ subsets: ['latin', 'cyrillic'] });

export const metadata: Metadata = {
  title: {
    default: 'Diaxel — AI‑ассистенты для автоматизации общения и поддержки',
    template: '%s — Diaxel'
  },
  description:
    'Платформа для создания, управления и анализа AI‑ассистентов: мультиканальная интеграция, настраиваемые инструкции и прозрачная аналитика.',
  metadataBase: new URL('https://example.com'),
  openGraph: {
    title: 'Diaxel — управление AI‑ассистентами',
    description:
      'Создавайте и масштабируйте AI‑ассистентов без глубоких знаний ML. Аналитика, чаты, интеграции.',
    type: 'website'
  },
  twitter: {
    card: 'summary_large_image'
  }
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="ru" className={inter.className}>
      <body>
        <div className="min-h-dvh bg-white text-slate-900">
          <Navbar />
          <main>{children}</main>
          <Footer />
        </div>
      </body>
    </html>
  );
}
