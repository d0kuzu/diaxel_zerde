import { clsx } from 'clsx';
import { forwardRef, type ComponentPropsWithoutRef, type ElementRef } from 'react';

export const Textarea = forwardRef<
  ElementRef<'textarea'>,
  ComponentPropsWithoutRef<'textarea'>
>(({ className, ...props }, ref) => {
  return (
    <textarea
      className={clsx(
        'flex min-h-[80px] w-full rounded-xl border border-slate-200 bg-white px-4 py-3 text-sm placeholder:text-slate-500 focus:border-brand-500 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:ring-offset-2 ring-offset-white disabled:cursor-not-allowed disabled:opacity-50',
        className
      )}
      ref={ref}
      {...props}
    />
  );
});

Textarea.displayName = 'Textarea';
