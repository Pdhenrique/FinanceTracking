export interface Transaction {
  id: string;
  agency: string;
  account_id: string;
  release_date: string;
  accounting_date: string;
  title: string;
  description: string;
  income: number;
  expense: number;
  daily_balance: number;
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