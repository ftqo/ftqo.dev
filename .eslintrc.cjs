/* eslint-env es6 */

/** @type {import('eslint').Linter.Config} */
module.exports = {
  extends: ['@remix-run/eslint-config/internal'],
  settings: {
    'import/internal-regex': '^~/',
  },
  plugins: ['prettier'],
  rules: {
    'import/order': [
      'error',
      {
        alphabetize: { caseInsensitive: true, order: 'asc' },
        groups: ['builtin', 'external', 'internal', 'parent', 'sibling'],
        'newlines-between': 'always',
      },
    ],

    'react/jsx-no-leaked-render': ['warn', { validStrategies: ['ternary'] }],
    'prettier/prettier': 'warn',
  },
}
