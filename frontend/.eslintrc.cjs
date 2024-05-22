/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  extends: [
    'eslint:recommended',
    'plugin:vue/vue3-recommended',
    'prettier',
    '@vue/eslint-config-prettier'
  ],
  plugins: [
    'vue',
    'prettier'
  ],
  parserOptions: {
    ecmaVersion: 'latest'
  },
  rules: {
    'tabWidth':4,
    'useTabs': false,
    'vue/multi-word-component-names': 0
  }

}
