import { clsx } from 'clsx';

export function Badge({
  children,
  className,
  variant = 'default'
}: {
  children: React.ReactNode;
  className?: string;
  variant?: 'default' | 'secondary' | 'outline';
}) {
  const variants = {
    default: 'inline-flex items-center rounded-full border border-brand-200 bg-brand-50 px-3 py-1 text-xs font-medium text-brand-800',
    secondary: 'inline-flex items-center rounded-full border border-slate-200 bg-slate-100 px-3 py-1 text-xs font-medium text-slate-700',
    outline: 'inline-flex items-center rounded-full border border-slate-300 bg-transparent px-3 py-1 text-xs font-medium text-slate-700'
  };

  return (
    <span className={clsx(variants[variant], className)}>
      {children}
    </span>
  );
}
