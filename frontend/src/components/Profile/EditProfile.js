import React, { useState } from 'react';
import API from '../../services/api';
import { useHistory } from 'react-router-dom';

const EditProfile = () => {
  const [bio, setBio] = useState('');
  const history = useHistory();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await API.put('/profile/edit', { bio });
      history.push('/profile');
    } catch (error) {
      console.error('Edit profile failed', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <textarea value={bio} onChange={(e) => setBio(e.target.value)} placeholder="Bio" required />
      <button type="submit">Save</button>
    </form>
  );
};

export default EditProfile;
