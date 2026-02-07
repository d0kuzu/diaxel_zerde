import { Card } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

export const metadata = {
  title: 'Контакты'
};

export default function ContactsPage() {
  return (
    <div className="mx-auto max-w-7xl container-px py-14">
      <div className="max-w-2xl">
        <h1 className="text-3xl font-semibold tracking-tight">Контакты</h1>
        <p className="mt-3 text-slate-600">
          Расскажите о задаче — поможем запустить AI‑ассистента, интеграции и аналитику.
        </p>
      </div>

      <div className="mt-10 grid grid-cols-1 gap-4 lg:grid-cols-3">
        <Card className="lg:col-span-2">
          <div className="text-sm font-semibold">Форма обратной связи</div>
          <div className="mt-5 grid grid-cols-1 gap-3 sm:grid-cols-2">
            <Input placeholder="Имя" name="name" />
            <Input placeholder="Email" name="email" type="email" />
          </div>
          <div className="mt-3">
            <Input placeholder="Тема" name="subject" />
          </div>
          <div className="mt-3">
            <textarea
              name="message"
              placeholder="Сообщение"
              className="min-h-32 w-full rounded-xl border border-slate-200 bg-white px-4 py-3 text-sm text-slate-900 shadow-sm outline-none transition focus:border-brand-300 focus:ring-2 focus:ring-brand-500/20"
            />
          </div>
          <div className="mt-5">
            <Button variant="primary" type="button">
              Отправить
            </Button>
          </div>
          <div className="mt-3 text-xs text-slate-500">
            Сейчас форма — демо UI. Можно подключить отправку через ваш API Gateway.
          </div>
        </Card>

        <Card>
          <div className="text-sm font-semibold">Данные</div>
          <div className="mt-3 text-sm text-slate-600">
            Email: hello@diaxel.ai
            <div className="mt-1">Телефон: +7 (000) 000‑00‑00</div>
          </div>

          <div className="mt-6 text-sm font-semibold">Соцсети</div>
          <div className="mt-3 flex flex-col gap-2 text-sm">
            <a className="text-slate-600 hover:text-slate-900" href="#">
              Telegram
            </a>
            <a className="text-slate-600 hover:text-slate-900" href="#">
              X (Twitter)
            </a>
            <a className="text-slate-600 hover:text-slate-900" href="#">
              LinkedIn
            </a>
          </div>

          <div className="mt-6 rounded-2xl border border-slate-200 bg-slate-50 p-4">
            <div className="text-xs font-medium text-slate-600">Карта</div>
            <div className="mt-2 text-sm text-slate-600">
              Вставьте сюда виджет карты (Yandex/Google) при необходимости.
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
}
