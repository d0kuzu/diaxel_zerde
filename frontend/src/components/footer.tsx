import Link from 'next/link';

export function Footer() {
  return (
    <footer className="border-t border-slate-200 bg-white">
      <div className="mx-auto max-w-7xl container-px py-10">
        <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-4">
          <div>
            <div className="text-sm font-semibold">Diaxel</div>
            <div className="mt-2 text-sm text-slate-600">
              Платформа управления AI‑ассистентами для автоматизации общения и поддержки.
            </div>
          </div>

          <div>
            <div className="text-sm font-semibold">Продукт</div>
            <div className="mt-3 flex flex-col gap-2 text-sm">
              <Link className="text-slate-600 hover:text-slate-900" href="/features">
                Функции
              </Link>
              <Link className="text-slate-600 hover:text-slate-900" href="/analytics">
                Аналитика
              </Link>
              <Link className="text-slate-600 hover:text-slate-900" href="/pricing">
                Тарифы
              </Link>
            </div>
          </div>

          <div>
            <div className="text-sm font-semibold">Ресурсы</div>
            <div className="mt-3 flex flex-col gap-2 text-sm">
              <Link className="text-slate-600 hover:text-slate-900" href="/blog">
                Блог
              </Link>
              <Link className="text-slate-600 hover:text-slate-900" href="/contacts">
                Контакты
              </Link>
            </div>
          </div>

          <div>
            <div className="text-sm font-semibold">Связь</div>
            <div className="mt-3 text-sm text-slate-600">
              hello@diaxel.ai
              <div className="mt-1">+7 (000) 000‑00‑00</div>
            </div>
          </div>
        </div>

        <div className="mt-10 flex flex-col gap-2 border-t border-slate-200 pt-6 text-xs text-slate-500 sm:flex-row sm:items-center sm:justify-between">
          <div>© {new Date().getFullYear()} Diaxel. Все права защищены.</div>
          <div className="flex gap-4">
            <a className="hover:text-slate-900" href="#">
              Privacy
            </a>
            <a className="hover:text-slate-900" href="#">
              Terms
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
}
