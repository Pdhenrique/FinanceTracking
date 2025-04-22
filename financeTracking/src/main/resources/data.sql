-- ---------- USERS ----------
INSERT INTO tb_users (id, name, email, password, cpf, active)
VALUES
 ('550e8400-e29b-41d4-a716-446655440000', 'Maria', 'maria@gmail.com', 'Teste12', '00000000000', true),
 ('660e8400-e29b-41d4-a716-556655440001', 'Bob'  , 'bob@gmail.com'  , 'este2' , '00000000001', true),
 ('770e8400-e29b-41d4-a716-666655440002', 'Alex' , 'alex@gmail.com' , 'dawd'  , '00000000002', true),
 ('880e8400-e29b-41d4-a716-776655440003', 'Ana'  , 'ana@gmail.com'  , 'awdwaaw','00000000003', true);

-- ---------- ACCOUNTS ----------
-- Conta corrente e poupança para cada user
INSERT INTO tb_accounts (id, bank_name, account_number, account_type, balance, active, user_id)
VALUES
 -- Maria
 ('a10e8400-e29b-41d4-a716-111155440000', 'Nubank' , '0001-1', 'CONTA_CORRENTE', 1500.00, true, '550e8400-e29b-41d4-a716-446655440000'),
 ('a20e8400-e29b-41d4-a716-111155440001', 'Nubank' , '0001-2', 'POUPANCA' ,  300.00, true, '550e8400-e29b-41d4-a716-446655440000'),

 -- Bob
 ('b10e8400-e29b-41d4-a716-222255440000', 'Inter'  , '0002-1', 'CONTA_CORRENTE',  950.00, true, '660e8400-e29b-41d4-a716-556655440001'),
 ('b20e8400-e29b-41d4-a716-222255440001', 'Inter'  , '0002-2', 'POUPANCA' , 1200.00, true, '660e8400-e29b-41d4-a716-556655440001'),

 -- Alex
 ('c10e8400-e29b-41d4-a716-333355440000', 'BB'     , '0003-1', 'CONTA_CORRENTE',  480.00, true, '770e8400-e29b-41d4-a716-666655440002'),
 ('c20e8400-e29b-41d4-a716-333355440001', 'BB'     , '0003-2', 'POUPANCA' ,   50.00, true, '770e8400-e29b-41d4-a716-666655440002'),

 -- Ana
 ('d10e8400-e29b-41d4-a716-444455440000', 'Caixa'  , '0004-1', 'CONTA_CORRENTE', 2300.00, true, '880e8400-e29b-41d4-a716-776655440003'),
 ('d20e8400-e29b-41d4-a716-444455440001', 'Caixa'  , '0004-2', 'POUPANCA' ,  400.00, true, '880e8400-e29b-41d4-a716-776655440003');
