import { clsx } from 'clsx';

import type { SelectHTMLAttributes } from 'react';

export function Select({
  className,
  ...props
}: SelectHTMLAttributes<HTMLSelectElement>) {
  return (
    <select
      {...props}
      className={clsx(
        'h-11 w-full rounded-xl border border-slate-200 bg-white px-4 text-sm text-slate-900 shadow-sm outline-none transition focus:border-brand-300 focus:ring-2 focus:ring-brand-500/20',
        className
      )}
    />
  );
}
