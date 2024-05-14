/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  extends: [
    'eslint:recommended',
    '@vue/eslint-config-prettier',
    'plugin:vue/vue3-essential',
    'prettier'
  ],
  plugins: [
    'vue',
    'prettier'
  ],
  parserOptions: {
    ecmaVersion: 'latest'
  }
}
