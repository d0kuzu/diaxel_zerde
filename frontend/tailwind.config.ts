import type { Config } from 'tailwindcss';

const config: Config = {
  content: ['./src/**/*.{js,ts,jsx,tsx,mdx}'],
  theme: {
    extend: {
      colors: {
        brand: {
          50: '#eff6ff',
          100: '#dbeafe',
          200: '#bfdbfe',
          300: '#93c5fd',
          400: '#60a5fa',
          500: '#3b82f6',
          600: '#2563eb',
          700: '#1d4ed8',
          800: '#1e40af',
          900: '#1e3a8a'
        },
        accent: {
          400: '#22d3ee',
          500: '#06b6d4',
          600: '#0891b2'
        }
      },
      boxShadow: {
        glow: '0 0 0 1px rgba(59,130,246,0.25), 0 10px 30px rgba(37,99,235,0.18)'
      }
    }
  },
  plugins: []
};

export default config;
