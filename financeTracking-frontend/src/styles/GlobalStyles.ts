import { createGlobalStyle } from 'styled-components';

export const GlobalStyles = createGlobalStyle`
  :root {
    /* Light theme (default) */
    --primary: #0047cc;
    --primary-light: #2563eb;
    --secondary: #00d4ff;
    --background: #f8fafc;
    --background-light: #ffffff;
    --text: #1e293b;
    --text-light: #64748b;
    --success: #10b981;
    --warning: #f59e0b;
    --danger: #ef4444;
    --border: #e2e8f0;
    --border-light: #f1f5f9;
    --shadow: rgba(0, 0, 0, 0.1);
  }
  
  /* Dark theme */
  [data-theme='dark'] {
    --primary: #3b82f6;
    --primary-light: #60a5fa;
    --secondary: #38bdf8;
    --background: #0f172a;
    --background-light: #1e293b;
    --text: #f1f5f9;
    --text-light: #94a3b8;
    --success: #34d399;
    --warning: #fbbf24;
    --danger: #f87171;
    --border: #334155;
    --border-light: #1e293b;
    --shadow: rgba(0, 0, 0, 0.3);
  }

  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    background-color: var(--background);
    color: var(--text);
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  button, input, select, textarea {
    font-family: inherit;
  }

  button {
    cursor: pointer;
  }

  a {
    color: inherit;
    text-decoration: none;
  }
`;

export default GlobalStyles;