/// <reference types="vitest" />
// typescriptにvitestの型定義を適用する、下段のtestがエラーになるのを防ぐ
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';
import tsconfigPaths from 'vite-tsconfig-paths';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  test: {
    globals: true,
    environment: 'happy-dom', // 高速に動作する
    setupFiles: ['./vitest-setup.ts'], // vitest-setup.tsを適用する、毎回importしなくて済む
  },
});
