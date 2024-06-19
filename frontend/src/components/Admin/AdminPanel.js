import React, { useEffect, useState } from 'react';
import API from '../../services/api';

const AdminPanel = () => {
  const [users, setUsers] = useState([]);
  const [dormantUsers, setDormantUsers] = useState([]);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await API.get('/admin/users');
        setUsers(response.data.users);
      } catch (error) {
        console.error('Error fetching users', error);
      }
    };

    const fetchDormantUsers = async () => {
      try {
        const response = await API.get('/admin/dormant-accounts');
        setDormantUsers(response.data.dormant_users);
      } catch (error) {
        console.error('Error fetching dormant accounts', error);
      }
    };

    fetchUsers();
    fetchDormantUsers();
  }, []);

  const deleteUser = async (userID) => {
    try {
      await API.delete(`/admin/user/${userID}`);
      setUsers(users.filter(user => user.id !== userID));
    } catch (error) {
      console.error('Error deleting user', error);
    }
  };

  return (
    <div>
      <h1>Admin Panel</h1>
      <h2>All Users</h2>
      <ul>
        {users.map(user => (
          <li key={user.id}>
            {user.name} ({user.email})
            <button onClick={() => deleteUser(user.id)}>Delete</button>
          </li>
        ))}
      </ul>
      <h2>Dormant Users</h2>
      <ul>
        {dormantUsers.map(user => (
          <li key={user.id}>
            {user.name} ({user.email})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AdminPanel;
