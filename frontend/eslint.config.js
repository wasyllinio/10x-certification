import js from '@eslint/js';
import tseslint from '@typescript-eslint/eslint-plugin';
import tsparser from '@typescript-eslint/parser';
import angular from '@angular-eslint/eslint-plugin';
import angularTemplateParser from '@angular-eslint/template-parser';
import prettier from 'eslint-config-prettier';

export default [
  // Base configuration for all files
  js.configs.recommended,
  
  // TypeScript files configuration
  {
    files: ['**/*.ts'],
    languageOptions: {
      parser: tsparser,
      parserOptions: {
        project: ['./tsconfig.app.json', './tsconfig.spec.json'],
        ecmaVersion: 2022,
        sourceType: 'module',
      },
      globals: {
        console: 'readonly',
      },
    },
    plugins: {
      '@typescript-eslint': tseslint,
      '@angular-eslint': angular,
    },
    rules: {
      // TypeScript recommended rules
      '@typescript-eslint/no-unused-vars': 'error',
      '@typescript-eslint/no-explicit-any': 'warn',
      '@typescript-eslint/explicit-function-return-type': 'off',
      '@typescript-eslint/explicit-module-boundary-types': 'off',
      '@typescript-eslint/no-empty-function': 'warn',
      
      // Angular recommended rules
      '@angular-eslint/directive-selector': 'error',
      '@angular-eslint/component-selector': 'error',
      '@angular-eslint/use-lifecycle-interface': 'error',
    },
  },
  
  // Test files configuration
  {
    files: ['**/*.spec.ts'],
    languageOptions: {
      parser: tsparser,
      parserOptions: {
        project: ['./tsconfig.app.json', './tsconfig.spec.json'],
        ecmaVersion: 2022,
        sourceType: 'module',
      },
      globals: {
        describe: 'readonly',
        it: 'readonly',
        expect: 'readonly',
        beforeEach: 'readonly',
        afterEach: 'readonly',
        jasmine: 'readonly',
        console: 'readonly',
      },
    },
    plugins: {
      '@typescript-eslint': tseslint,
      '@angular-eslint': angular,
    },
    rules: {
      // TypeScript recommended rules
      '@typescript-eslint/no-unused-vars': 'error',
      '@typescript-eslint/no-explicit-any': 'warn',
      '@typescript-eslint/explicit-function-return-type': 'off',
      '@typescript-eslint/explicit-module-boundary-types': 'off',
      '@typescript-eslint/no-empty-function': 'warn',
      
      // Angular recommended rules
      '@angular-eslint/directive-selector': 'error',
      '@angular-eslint/component-selector': 'error',
      '@angular-eslint/use-lifecycle-interface': 'error',
    },
  },
  
  // Angular template files configuration
  {
    files: ['**/*.html'],
    languageOptions: {
      parser: angularTemplateParser,
      parserOptions: {
        project: ['./tsconfig.app.json', './tsconfig.spec.json'],
        ecmaVersion: 2022,
        sourceType: 'module',
      },
    },
    plugins: {
      '@angular-eslint': angular,
    },
    rules: {
      // Basic template rules - keeping it simple for now
    },
  },
  
  // Ignore patterns
  {
    ignores: [
      'node_modules/**',
      'dist/**',
      '.angular/**',
      'coverage/**',
      '*.config.js',
      '*.config.ts',
    ],
  },
  
  // Prettier integration (must be last)
  prettier,
];
