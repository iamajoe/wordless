module.exports = {
  env: {
    es2021: true,
  },
  extends: [
    'plugin:react/recommended',
    'airbnb',
    'plugin:@typescript-eslint/recommended',
    'prettier',
    'prettier/prettier',
    'plugin:prettier/recommended',
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 12,
    sourceType: 'module',
  },
  plugins: ['react', '@typescript-eslint'],
  rules: {
    'no-continue': 'off',
    'no-nested-ternary': 'off',
    'no-throw-literal': 'off',
    'no-case-declarations': 'off',
    'no-async-promise-executor': 'off',
    'prefer-promise-reject-errors': 'warn',
    'no-loop-func': 'off',
    'no-await-in-loop': 'off',
    'lines-between-class-members': [
      'error',
      'always',
      { exceptAfterSingleLine: true },
    ],

    'import/no-extraneous-dependencies': ['error', {}],
    'no-underscore-dangle': 'off',
    'no-use-before-define': 'off',
    '@typescript-eslint/no-use-before-define': ['error'],
    '@typescript-eslint/ban-types': 'off',
    'import/extensions': [
      'error',
      'ignorePackages',
      {
        ts: 'never',
        tsx: 'never',
      },
    ],
    'no-shadow': 'off',
    '@typescript-eslint/no-shadow': ['error'],
    // '@typescript-eslint/explicit-function-return-type': [
    //   'error',
    //   {
    //     allowExpressions: true,
    //   },
    // ],
    'max-len': [
      'warn',
      {
        code: 110,
      },
    ],
    'import/prefer-default-export': 'off',
  },
  settings: {
    'import/resolver': {
      typescript: {},
    },
  },
};
