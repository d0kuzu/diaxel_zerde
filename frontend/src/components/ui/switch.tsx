'use client';

import { clsx } from 'clsx';
import { forwardRef, type ComponentPropsWithoutRef, type ElementRef } from 'react';

export const Switch = forwardRef<
  ElementRef<'button'>,
  ComponentPropsWithoutRef<'button'> & {
    checked?: boolean;
    onCheckedChange?: (checked: boolean) => void;
  }
>(({ className, checked, onCheckedChange, onClick, ...props }, ref) => {
  const handleClick = (e: React.MouseEvent<HTMLButtonElement>) => {
    onCheckedChange?.(!checked);
    onClick?.(e);
  };

  return (
    <button
      type="button"
      role="switch"
      aria-checked={checked}
      ref={ref}
      className={clsx(
        'peer inline-flex h-6 w-11 shrink-0 cursor-pointer items-center rounded-full border-2 border-transparent transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-brand-500 focus-visible:ring-offset-2 focus-visible:ring-offset-white disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-brand-600 data-[state=unchecked]:bg-slate-200',
        checked && 'data-[state=checked]',
        !checked && 'data-[state=unchecked]',
        className
      )}
      onClick={handleClick}
      data-state={checked ? 'checked' : 'unchecked'}
      {...props}
    >
      <span
        data-state={checked ? 'checked' : 'unchecked'}
        className={clsx(
          'pointer-events-none block h-5 w-5 rounded-full bg-white shadow-lg ring-0 transition-transform data-[state=checked]:translate-x-5 data-[state=unchecked]:translate-x-0',
          checked && 'data-[state=checked]',
          !checked && 'data-[state=unchecked]'
        )}
      />
    </button>
  );
});

Switch.displayName = 'Switch';
