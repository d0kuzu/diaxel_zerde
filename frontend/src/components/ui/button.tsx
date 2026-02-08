import Link from 'next/link';

import type { ButtonHTMLAttributes, ReactElement, ReactNode } from 'react';

type CommonProps = {
  children: ReactNode;
  className?: string;
  variant?: 'primary' | 'secondary' | 'ghost' | 'outline';
  size?: 'sm' | 'default' | 'lg';
};

type ButtonAsButton = CommonProps &
  ButtonHTMLAttributes<HTMLButtonElement> & { asChild?: false };

type ButtonAsChild = CommonProps & { asChild: true; href?: never };

export function Button(props: ButtonAsButton | ButtonAsChild) {
  const { className, children } = props;
  const variant: NonNullable<CommonProps['variant']> = props.variant ?? 'primary';
  const size: NonNullable<CommonProps['size']> = props.size ?? 'default';

  const base =
    'inline-flex items-center justify-center rounded-xl font-medium transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-brand-500 focus-visible:ring-offset-2 ring-offset-white disabled:opacity-50 disabled:pointer-events-none';

  const sizes = {
    sm: 'px-3 py-2 text-xs',
    default: 'px-5 py-3 text-sm',
    lg: 'px-6 py-4 text-base'
  };

  const stylesByVariant: Record<NonNullable<CommonProps['variant']>, string> = {
    primary:
      'bg-brand-600 text-white shadow-glow hover:bg-brand-700 active:bg-brand-800',
    secondary:
      'bg-white text-slate-900 border border-slate-200 hover:bg-slate-50 active:bg-slate-100',
    ghost: 'bg-transparent text-slate-900 hover:bg-slate-100 active:bg-slate-200',
    outline: 'bg-transparent text-slate-900 border border-slate-200 hover:bg-slate-50 active:bg-slate-100'
  };

  const styles = [base, sizes[size], stylesByVariant[variant], className].filter(Boolean).join(' ');

  if ('asChild' in props && props.asChild) {
    const child = children as ReactElement<{ href: string; children: ReactNode }>;
    return (
      <Link
        href={child.props.href}
        className={styles}
      >
        {child.props.children}
      </Link>
    );
  }

  const { asChild, ...rest } = props;

  return (
    <button {...rest} className={styles}>
      {children}
    </button>
  );
}
