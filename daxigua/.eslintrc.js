module.exports = {
  env: {
    browser: true,
    node: true,
    commonjs: true,
    amd: true,
    es6: true,
    mocha: true,
  },
  extends: [
    'eslint:recommended'
  ],
  parser: 'babel-eslint',
  parserOptions: {
    ecmaVersion: 7,
    sourceType: 'module',
  },
  rules: {
    'no-nested-ternary': 'warn',
    'max-classes-per-file': 'off',
    'import/no-unresolved': 'off',
    'import/extensions': 'off',
    'no-trailing-spaces': 'off',
    'no-bitwise': 'off',
    'no-multi-spaces': 'off',
    'spaced-comment': 'off',
    'no-confusing-arrow': ['error', { allowParens: true }],
    'no-plusplus': 'off',
    'prefer-template': 'warn',
    'import/no-dynamic-require': 'off',
    'import/prefer-default-export': 'off',
    'global-require': 'off',
    'no-return-assign': ['error', 'except-parens'],
    'no-unused-expressions': [
      'error',
      {
        allowShortCircuit: true,
        allowTernary: true,
      }
    ],
    'arrow-body-style': 'off',
    'class-methods-use-this': 'off',
    'no-restricted-syntax': [
      'error',
      {
        selector: 'ForInStatement',
        message: 'for..in loops iterate over the entire prototype chain, which is virtually never what you want. Use Object.{keys,values,entries}, and iterate over the resulting array.',
      },
      {
        selector: 'LabeledStatement',
        message: 'Labels are a form of GOTO; using them makes code confusing and hard to maintain and understand.',
      },
      {
        selector: 'WithStatement',
        message: '`with` is disallowed in strict mode because it makes code impossible to predict and optimize.',
      }
    ],
    'no-unused-vars': [
      'error',
      {
        varsIgnorePattern: '[Ll]ogger',
        ignoreRestSiblings: true,
      }
    ],
    'no-magic-numbers': [
      'warn',
      {
        enforceConst: true,
        ignore: [-1, 0, 1],
        ignoreArrayIndexes: true,
      }
    ],
    'no-warning-comments': [
      'warn',
      {
        terms: ['todo', 'fixme'],
        location: 'anywhere',
      }
    ],
    camelcase: [
      'error',
      {
        properties: 'never',
      }
    ],
    'func-names': ['error', 'as-needed'],
    'comma-dangle': [
      'error',
      {
        functions: 'never',
        arrays: 'never',
        imports: 'never',
        exports: 'never',
        objects: 'always-multiline',
      }
    ],
  },
};

module.rules = {
  'linebreak-style': [0, 'error', 'windows'],
  'no-unused-vars': 'off',
};
