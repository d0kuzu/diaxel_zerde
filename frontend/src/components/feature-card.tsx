export function FeatureCard({
  title,
  description
}: {
  title: string;
  description: string;
}) {
  return (
    <div className="group rounded-2xl border border-slate-200 bg-white p-6 shadow-sm transition-all hover:-translate-y-0.5 hover:shadow-md">
      <div className="flex items-start justify-between gap-4">
        <div>
          <div className="text-sm font-semibold tracking-tight">{title}</div>
          <div className="mt-2 text-sm text-slate-600">{description}</div>
        </div>
        <div className="h-10 w-10 rounded-xl bg-brand-50 ring-1 ring-brand-100 transition-colors group-hover:bg-brand-100" />
      </div>
    </div>
  );
}
