# Transações no Front‑end

Este documento explica como a funcionalidade de transações funciona no front‑end: modelo de dados, fluxo, hook (`useTransactions`), componentes que consomem os dados e pontos de atenção para evoluções.

## Visão Geral

- Os dados de transações são buscados via API e armazenados/gerenciados com React Query.
- Resumos (saldo e categorias) são calculados no cliente a partir dos campos `income` e `expense`.
- A dashboard exibe:
  - Lista de transações (com ícones por tipo derivado de income/expense)
  - Gráfico de despesas por categoria (título)
  - Gráfico de evolução (entradas x saídas por dia)

## Modelo de Dados

A interface principal está em `src/types/index.ts:1`:

```ts
export interface Transaction {
  id: string;
  agency: string;
  account_id: string;
  release_date: string; // ISO: "2024-08-16T00:00:00Z"
  accounting_date: string; // ISO
  title: string; // usado como "categoria"
  description: string;
  income: number; // valor de entrada
  expense: number; // valor de saída
  daily_balance: number;
}
```

Resumos calculados no front:

```ts
export interface BalanceSummary {
  totalBalance: number;
  income: number;
  expenses: number;
}
export interface CategorySummary {
  category: string;
  amount: number;
  percentage: number;
}
```

## Fluxo de Dados

1. Busca: `GET http://localhost:8080/v1/transactions` retorna `Transaction[]`.
2. Cache: React Query armazena e invalida o cache quando `add`/`remove` ocorre.
3. Derivação:
   - `BalanceSummary`: soma de `income` e `expense` e diferença.
   - `CategorySummary`: agrupa somente despesas por `title` e calcula percentuais.

## Hook useTransactions

Arquivo: `src/hooks/useTransactions.ts:26`

Retorno principal:

```ts
{
  // dados
  transactions: Transaction[],
  balanceSummary: BalanceSummary,
  categorySummary: CategorySummary[],

  // ações
  addTransaction: (tx: Omit<Transaction, 'id'>) => void,
  removeTransaction: (id: string) => void,

  // estados
  isLoading: boolean,
  isError: boolean,
  isAddingTransaction: boolean,
  isRemovingTransaction: boolean,
  addTransactionError?: unknown,
  removeTransactionError?: unknown,
}
```

Como funciona:

- Busca: `fetchTransactions()` (`GET /v1/transactions`).
- Adição: `addTransactionAPI(tx)` (`POST /v1/transactions`).
- Remoção: `removeTransactionAPI(id)` (`DELETE /v1/transactions/{id}`).
- Resumos são recalculados via `useEffect` sempre que `transactions` muda.

## Componentes

- Dashboard (`src/components/dashboard/Dashboard.tsx:10`)

  - Consome o hook e passa dados para os componentes filhos.

- Lista de Transações (`src/components/transactions/TransactionList.tsx:64`)

  - Exibe `description`, `title` (como categoria), `release_date` (formatada) e valor.
  - O tipo é derivado: `expense > 0` → "expense"; `income > 0` → "income"; caso contrário "transfer".

- Gráfico de Despesas por Categoria (`src/components/charts/ExpensesChart.tsx`)

  - Entrada: `CategorySummary[]`.

- Tendência de Transações (`src/components/charts/TransactionTrendsChart.tsx:14`)

  - Agrupa por data usando `release_date` (fallback `accounting_date`).
  - Soma `income` e `expense` por dia e ordena pelo eixo X.

## Estados e Erros

- `isLoading` e `isError` vêm do React Query.
- `isAddingTransaction`, `isRemovingTransaction` e erros associados vêm das mutations.
- O cache de `transactions` é invalidado após `add` e `remove` para sincronizar a UI.

## Alinhamento com o Back‑end (Importante)

- O backend trabalha com `income` e `expense` (separados) e datas `release_date`/`accounting_date`.
- Alguns componentes legados ainda usam o formato antigo (`amount`, `type`, `date`, `category`).
  - Ex.: `AddTransactionForm` monta `{ description, amount, category, type, date }`.
  - Para enviar ao backend, é necessário mapear para o modelo atual:

```ts
// exemplo de transformação antes do POST
const payload: Omit<Transaction, "id"> = {
  agency: "1", // definir conforme necessidade
  account_id: "...", // idem
  release_date: date, // do input
  accounting_date: date, // ou outra data
  title: category,
  description,
  income: type === "income" ? Number(amount) : 0,
  expense: type === "expense" ? Number(amount) : 0,
  daily_balance: 0, // opcionalmente calculado no backend
};
```

> Recomendação: ajustar `AddTransactionForm` para construir o payload acima e usar `addTransaction` do hook. Se quiser, há como mover essa transformação para dentro do próprio hook.

## Exemplos de Uso

Uso básico na dashboard:

```tsx
const {
  transactions,
  balanceSummary,
  categorySummary,
  addTransaction,
  removeTransaction,
  isLoading,
  isError,
} = useTransactions();

if (isLoading) return <Spinner />;
if (isError) return <Error />;

return (
  <>
    <BalanceCard {...balanceSummary} />
    <ExpensesChart data={categorySummary} />
    <TransactionTrendsChart transactions={transactions} />
    <TransactionList
      transactions={transactions}
      onRemoveTransaction={removeTransaction}
    />
  </>
);
```

## Boas Práticas e Extensões

- Validação: Garantir que somente um de `income`/`expense` seja positivo por transação.
- Data: Preferir `release_date` para exibição; cair para `accounting_date` quando necessário.
- Filtros: Adicionar filtros por intervalo de datas e por `title` antes de calcular resumos.
- Paginação: Se a API suportar, integrar via `useInfiniteQuery` para listas grandes.
- Internacionalização: A formatação usa `Intl` com locale `pt-BR` para valores e datas.

---

Em caso de dúvidas ou para evoluir o formulário de criação para o modelo do backend, abra uma issue ou peça para ajustarmos o `AddTransactionForm` e/ou o próprio hook.
