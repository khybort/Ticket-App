import React from 'react';
import { Link } from 'react-router-dom';
import Navbar from '../components/Navbar/Navbar';

const Home: React.FC = () => {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <div className="container mx-auto p-8 text-center">
        <h1 className="text-4xl font-bold text-blue-600 mb-6">
          Welcome to the Ticketing System
        </h1>
        <p className="text-lg text-gray-700 mb-8">
          Manage and purchase your tickets easily.
        </p>
        <div className="space-x-4">
          <Link
            to="/ticket-list"
            className="text-white bg-blue-500 px-6 py-2 rounded-lg hover:bg-blue-600"
          >
            View Tickets
          </Link>
          <Link
            to="/create-ticket"
            className="text-white bg-green-500 px-6 py-2 rounded-lg hover:bg-green-600"
          >
            Create Ticket
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Home;
