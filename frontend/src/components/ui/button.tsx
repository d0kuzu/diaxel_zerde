import Link from 'next/link';
import { clsx } from 'clsx';

import type { ButtonHTMLAttributes, ReactElement, ReactNode } from 'react';

type CommonProps = {
  children: ReactNode;
  className?: string;
  variant?: 'primary' | 'secondary' | 'ghost';
};

type ButtonAsButton = CommonProps &
  ButtonHTMLAttributes<HTMLButtonElement> & { asChild?: false };

type ButtonAsChild = CommonProps & { asChild: true; href?: never };

export function Button(props: ButtonAsButton | ButtonAsChild) {
  const { className, children } = props;
  const variant: NonNullable<CommonProps['variant']> = props.variant ?? 'primary';

  const base =
    'inline-flex items-center justify-center rounded-xl px-5 py-3 text-sm font-medium transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-brand-500 focus-visible:ring-offset-2 ring-offset-white disabled:opacity-50 disabled:pointer-events-none';

  const stylesByVariant: Record<NonNullable<CommonProps['variant']>, string> = {
    primary:
      'bg-brand-600 text-white shadow-glow hover:bg-brand-700 active:bg-brand-800',
    secondary:
      'bg-white text-slate-900 border border-slate-200 hover:bg-slate-50 active:bg-slate-100',
    ghost: 'bg-transparent text-slate-900 hover:bg-slate-100 active:bg-slate-200'
  };

  const styles = stylesByVariant[variant];

  if ('asChild' in props && props.asChild) {
    const child = children as ReactElement<{ href: string; children: ReactNode }>;
    return (
      <Link
        href={child.props.href}
        className={clsx(base, styles, className)}
      >
        {child.props.children}
      </Link>
    );
  }

  const { asChild, ...rest } = props;

  return (
    <button {...rest} className={clsx(base, styles, className)}>
      {children}
    </button>
  );
}
