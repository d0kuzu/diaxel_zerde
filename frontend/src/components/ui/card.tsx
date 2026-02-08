import { clsx } from 'clsx';

import type { ReactNode } from 'react';

export function Card({
  className,
  children
}: {
  className?: string;
  children: ReactNode;
}) {
  return (
    <div className={clsx('rounded-2xl border border-slate-200 bg-white p-6 shadow-sm', className)}>
      {children}
    </div>
  );
}
