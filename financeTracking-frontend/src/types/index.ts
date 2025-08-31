export type TransactionType = 'income' | 'expense' | 'transfer';

export interface Transaction {
  id: string;
  description: string;
  amount: number;
  date: string;
  category: string;
  type: TransactionType;
}

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