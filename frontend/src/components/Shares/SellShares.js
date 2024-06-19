import React, { useState } from 'react';
import API from '../../services/api';
import { useHistory } from 'react-router-dom';

const SellShares = () => {
  const [amount, setAmount] = useState(0);
  const [days, setDays] = useState(0);
  const history = useHistory();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await API.post('/shares/sell', { amount, days });
      history.push('/shares');
    } catch (error) {
      console.error('Error selling shares', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Amount:</label>
        <input 
          type="number" 
          value={amount} 
          onChange={(e) => setAmount(e.target.value)} 
          placeholder="Amount" 
          required 
        />
      </div>
      <div>
        <label>Days:</label>
        <input 
          type="number" 
          value={days} 
          onChange={(e) => setDays(e.target.value)} 
          placeholder="Days" 
          required 
        />
      </div>
      <button type="submit">Sell Shares</button>
    </form>
  );
};

export default SellShares;
