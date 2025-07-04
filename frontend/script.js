const API_URL = 'https://financial-tool-web.onrender.com/api/transactions'; 

document.addEventListener('DOMContentLoaded', () => {
  fetchGoals();

  document.getElementById('goal-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    const title = document.getElementById('goal-title').value;
    const target = parseFloat(document.getElementById('goal-target').value);
    const current = parseFloat(document.getElementById('goal-current').value);

    await fetch('https://financial-tool-web.onrender.com/api/goals', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        title,
        target_amount: target,
        current_amount: current
      }),
    });

    e.target.reset();
    fetchGoals();
  });


  fetchTransactions();

  document.getElementById('transaction-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    const category = document.getElementById('category').value;
    const title = document.getElementById('title').value;
    const amount = parseFloat(document.getElementById('amount').value);
    const type = document.getElementById('type').value;

    if (!title || isNaN(amount)) return;

    await fetch(API_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        category,
        title,
        amount,
        type
      }),
    });

    e.target.reset();
    fetchTransactions();
  });
});

async function fetchTransactions() {
  const res = await fetch(API_URL);
  const data = await res.json();

  const list = document.getElementById('transactions-list');
  list.innerHTML = '';

  let income = 0;
  let expense = 0;

  data.forEach(t => {
    const li = document.createElement('li');
    li.innerHTML = `
      <span class="${t.type === 'income' ? 'text-green' : 'text-red'}">
       ${t.category.toUpperCase()} | ${t.title} | ${t.type === 'income' ? '+' : '–'}$${t.amount.toFixed(2)} (${t.type})
      </span>
      <button class="delete-btn" onclick="deleteTransaction(${t.id})">Delete</button>
    `;
    list.appendChild(li);

    if (t.type === 'income') income += t.amount;
    else expense += t.amount;
  });

  const balance = income - expense;
  document.getElementById('balance').textContent = `$${balance.toFixed(2)}`;
  document.getElementById('income').textContent = `$${income.toFixed(2)}`;
  document.getElementById('expenses').textContent = `$${expense.toFixed(2)}`;
}

async function deleteTransaction(id) {
  await fetch(`${API_URL}/${id}`, {
    method: 'DELETE'
  });
  await fetchTransactions();
}

async function fetchGoals() {
  const res = await fetch('https://financial-tool-web.onrender.com/api/goals');
  const goals = await res.json();

  const list = document.getElementById('goals-list');
  list.innerHTML = '';

  goals.forEach(goal => {
    const percent = Math.min((goal.current_amount / goal.target_amount) * 100, 100);

    const li = document.createElement('li');
    li.className = 'goal-item';
    li.innerHTML = `
      <div class="goal-title">${goal.title.toUpperCase()} — $${goal.current_amount.toFixed(2)} / $${goal.target_amount.toFixed(2)}</div>
      <div class="goal-progress">
        <div class="goal-progress-fill" style="width: ${percent}%"></div>
      </div>
      <input type="number" id="goal-input-${goal.id}" placeholder="Amount" style="margin-right: 20px;" />
      <button class="goal-delete-btn" onclick="fulfillGoal(${goal.id})">Add</button>
      <button class="goal-delete-btn" onclick="minusFromGoal(${goal.id})">Minus</button>
      <button class="goal-delete-btn" onclick="deleteGoal(${goal.id})">Delete</button>
    `;

    list.appendChild(li);
  });
}

async function deleteGoal(id) {
  await fetch(`https://financial-tool-web.onrender.com/api/goals/${id}`, { method: 'DELETE' });
  await fetchGoals();
}

async function fulfillGoal(id) {
  const input = document.getElementById(`goal-input-${id}`);
  const amount = parseFloat(input.value);
  if (isNaN(amount) || amount <= 0) return;

  await fetch(`https://financial-tool-web.onrender.com/api/goals/${id}/add`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ amount })
  });

  input.value = '';
  await fetchGoals();
}

async function minusFromGoal(id) {
  const input = document.getElementById(`goal-input-${id}`);
  const amount = parseFloat(input.value);
  if (isNaN(amount) || amount <= 0) return;

  await fetch(`https://financial-tool-web.onrender.com/api/goals/${id}/minus`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ amount })
  });

  input.value = '';
  await fetchGoals();
}

