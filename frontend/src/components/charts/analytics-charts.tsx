'use client';

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Tooltip,
  Legend
} from 'chart.js';
import { Line } from 'react-chartjs-2';
import type { AnalyticsTimeseriesPoint } from '@/lib/api';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Tooltip, Legend);

export function AnalyticsCharts({ data }: { data: AnalyticsTimeseriesPoint[] }) {
  const labels = data.map((p) => p.date.slice(5));

  const chats = data.map((p) => p.chats);
  const users = data.map((p) => p.activeUsers);
  const er = data.map((p) => Math.round(p.engagementRate * 100));

  return (
    <div className="grid grid-cols-1 gap-4 lg:grid-cols-2">
      <div className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
        <div className="text-sm font-semibold">Чаты</div>
        <div className="mt-4 h-64">
          <Line
            data={{
              labels,
              datasets: [
                {
                  label: 'Чаты',
                  data: chats,
                  borderColor: 'rgb(37, 99, 235)',
                  backgroundColor: 'rgba(37, 99, 235, 0.15)',
                  tension: 0.35
                }
              ]
            }}
            options={{
              responsive: true,
              maintainAspectRatio: false,
              plugins: {
                legend: { display: false }
              },
              scales: {
                y: { grid: { color: 'rgba(15,23,42,0.06)' } },
                x: { grid: { display: false } }
              }
            }}
          />
        </div>
      </div>

      <div className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
        <div className="text-sm font-semibold">Активные пользователи</div>
        <div className="mt-4 h-64">
          <Line
            data={{
              labels,
              datasets: [
                {
                  label: 'Пользователи',
                  data: users,
                  borderColor: 'rgb(6, 182, 212)',
                  backgroundColor: 'rgba(6, 182, 212, 0.18)',
                  tension: 0.35
                }
              ]
            }}
            options={{
              responsive: true,
              maintainAspectRatio: false,
              plugins: {
                legend: { display: false }
              },
              scales: {
                y: { grid: { color: 'rgba(15,23,42,0.06)' } },
                x: { grid: { display: false } }
              }
            }}
          />
        </div>
      </div>

      <div className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm lg:col-span-2">
        <div className="text-sm font-semibold">Engagement rate, %</div>
        <div className="mt-4 h-64">
          <Line
            data={{
              labels,
              datasets: [
                {
                  label: 'ER',
                  data: er,
                  borderColor: 'rgb(59, 130, 246)',
                  backgroundColor: 'rgba(59, 130, 246, 0.14)',
                  tension: 0.35
                }
              ]
            }}
            options={{
              responsive: true,
              maintainAspectRatio: false,
              plugins: {
                legend: { display: false }
              },
              scales: {
                y: {
                  suggestedMin: 0,
                  suggestedMax: 100,
                  grid: { color: 'rgba(15,23,42,0.06)' }
                },
                x: { grid: { display: false } }
              }
            }}
          />
        </div>
      </div>
    </div>
  );
}
