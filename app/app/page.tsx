'use client'
import React, {useEffect} from 'react';
import Login from './Login'; // Adjust the import path as needed

export default function Dashboard() {
  useEffect(() => {
    // Mark this component as a client entry point
  }, []);
  return (
    <div>
      <h1>Dashboard</h1>
      <p>Welcome to the dashboard!</p>
      
      {/* Render the Login component */}
      <Login />
    </div>
  );
}