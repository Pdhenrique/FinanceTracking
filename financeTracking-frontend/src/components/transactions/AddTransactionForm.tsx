import React, { useState } from 'react';
import styled from 'styled-components';
import type { Transaction, TransactionType } from '../../types';

interface AddTransactionFormProps {
  onAddTransaction: (transaction: Omit<Transaction, 'id'>) => void;
}

const AddTransactionForm: React.FC<AddTransactionFormProps> = ({ onAddTransaction }) => {
  const [description, setDescription] = useState('');
  const [amount, setAmount] = useState('');
  const [category, setCategory] = useState('');
  const [type, setType] = useState<TransactionType>('expense');
  const [date, setDate] = useState(new Date().toISOString().substring(0, 10));
  const [isFormOpen, setIsFormOpen] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (!description || !amount || !category) {
      alert('Por favor, preencha todos os campos.');
      return;
    }

    const newTransaction = {
      description,
      amount: Number(amount),
      category,
      type,
      date,
    };

    onAddTransaction(newTransaction);
    resetForm();
    setIsFormOpen(false);
  };

  const resetForm = () => {
    setDescription('');
    setAmount('');
    setCategory('');
    setType('expense');
    setDate(new Date().toISOString().substring(0, 10));
  };

  return (
    <FormContainer>
      {!isFormOpen ? (
        <AddButton onClick={() => setIsFormOpen(true)}>
          + Nova Transação
        </AddButton>
      ) : (
        <FormCard>
          <FormHeader>
            <FormTitle>Nova Transação</FormTitle>
            <CloseButton onClick={() => setIsFormOpen(false)}>×</CloseButton>
          </FormHeader>

          <Form onSubmit={handleSubmit}>
            <FormGroup>
              <Label>Descrição</Label>
              <Input
                type="text"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Ex: Salário, Aluguel, etc."
                required
              />
            </FormGroup>

            <FormGroup>
              <Label>Valor</Label>
              <Input
                type="number"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                placeholder="0,00"
                step="0.01"
                min="0"
                required
              />
            </FormGroup>

            <FormGroup>
              <Label>Categoria</Label>
              <Input
                type="text"
                value={category}
                onChange={(e) => setCategory(e.target.value)}
                placeholder="Ex: Alimentação, Transporte, etc."
                required
              />
            </FormGroup>

            <FormGroup>
              <Label>Tipo</Label>
              <TypeSelector>
                <TypeOption
                  selected={type === 'income'}
                  onClick={() => setType('income')}
                  type="income"
                >
                  Entrada
                </TypeOption>
                <TypeOption
                  selected={type === 'expense'}
                  onClick={() => setType('expense')}
                  type="expense"
                >
                  Saída
                </TypeOption>
                <TypeOption
                  selected={type === 'transfer'}
                  onClick={() => setType('transfer')}
                  type="transfer"
                >
                  Transferência
                </TypeOption>
              </TypeSelector>
            </FormGroup>

            <FormGroup>
              <Label>Data</Label>
              <Input
                type="date"
                value={date}
                onChange={(e) => setDate(e.target.value)}
                required
              />
            </FormGroup>

            <ButtonGroup>
              <CancelButton type="button" onClick={() => setIsFormOpen(false)}>
                Cancelar
              </CancelButton>
              <SubmitButton type="submit">
                Adicionar
              </SubmitButton>
            </ButtonGroup>
          </Form>
        </FormCard>
      )}
    </FormContainer>
  );
};

const FormContainer = styled.div`
  margin-bottom: 24px;
`;

const AddButton = styled.button`
  background-color: var(--primary);
  color: white;
  border: none;
  border-radius: 8px;
  padding: 12px 24px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s ease;
  width: 100%;
  
  &:hover {
    background-color: var(--primary-light);
  }
`;

const FormCard = styled.div`
  background-color: var(--background-light);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px var(--shadow);
`;

const FormHeader = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
`;

const FormTitle = styled.h3`
  font-size: 18px;
  font-weight: 600;
  color: var(--text);
  margin: 0;
`;

const CloseButton = styled.button`
  background: transparent;
  border: none;
  color: var(--text-light);
  font-size: 24px;
  cursor: pointer;
  line-height: 1;
  padding: 0;
  
  &:hover {
    color: var(--text);
  }
`;

const Form = styled.form`
  display: flex;
  flex-direction: column;
  gap: 16px;
`;

const FormGroup = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const Label = styled.label`
  font-size: 14px;
  font-weight: 500;
  color: var(--text);
`;

const Input = styled.input`
  padding: 12px 16px;
  border: 1px solid var(--border);
  border-radius: 8px;
  font-size: 14px;
  color: var(--text);
  background-color: var(--background);
  transition: border-color 0.2s ease;
  
  &:focus {
    outline: none;
    border-color: var(--primary);
  }
  
  &::placeholder {
    color: var(--text-light);
  }
`;

const TypeSelector = styled.div`
  display: flex;
  gap: 8px;
`;

const TypeOption = styled.button<{ selected: boolean; type: string }>`
  flex: 1;
  padding: 12px;
  border: 1px solid ${(props) => (props.selected ? 'transparent' : 'var(--border)')};
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  
  ${({ selected, type }) => {
    if (selected) {
      switch (type) {
        case 'income':
          return `
            background-color: var(--success);
            color: white;
          `;
        case 'expense':
          return `
            background-color: var(--danger);
            color: white;
          `;
        case 'transfer':
          return `
            background-color: var(--warning);
            color: white;
          `;
        default:
          return '';
      }
    } else {
      return `
        background-color: var(--background);
        color: var(--text);
        
        &:hover {
          background-color: var(--border-light);
        }
      `;
    }
  }}
`;

const ButtonGroup = styled.div`
  display: flex;
  gap: 12px;
  margin-top: 8px;
`;

const CancelButton = styled.button`
  flex: 1;
  padding: 12px;
  background-color: var(--background);
  border: 1px solid var(--border);
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  color: var(--text);
  cursor: pointer;
  transition: all 0.2s ease;
  
  &:hover {
    background-color: var(--border-light);
  }
`;

const SubmitButton = styled.button`
  flex: 1;
  padding: 12px;
  background-color: var(--primary);
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  color: white;
  cursor: pointer;
  transition: background-color 0.2s ease;
  
  &:hover {
    background-color: var(--primary-light);
  }
`;

export default AddTransactionForm;