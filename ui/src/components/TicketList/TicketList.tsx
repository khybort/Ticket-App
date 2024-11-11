import React from 'react';
import TicketItem from './TicketItem';
import TicketModal from './TicketModal';
import UserInput from '../UserInput/UserInput';
import useTicketState from '../../hooks/useTicketState';

const TicketList: React.FC = () => {
  const {
    tickets,
    loading,
    error,
    userId,
    quantities,
    selectedTicket,
    isModalOpen,
    handlePurchase,
    handleQuantityChange,
    handleGetTicket,
    closeModal,
    setUserId,
  } = useTicketState();

  return (
    <div className="max-w-4xl mx-auto p-6">
      <h2 className="text-3xl font-bold text-center mb-6">Available Tickets</h2>

      {error && <p className="text-red-500 text-center mb-4">{error}</p>}

      <UserInput userId={userId} onChange={setUserId} />

      <div className="bg-white shadow-lg rounded-lg p-6">
        <ul className="space-y-4">
          {tickets.map((ticket) => (
            <TicketItem
              key={ticket.id}
              ticket={ticket}
              onGetTicket={handleGetTicket}
              quantities={quantities}
              onQuantityChange={handleQuantityChange}
              onPurchase={handlePurchase}
              loading={loading}
            />
          ))}
        </ul>
      </div>

      {isModalOpen && selectedTicket && (
        <TicketModal ticket={selectedTicket} onClose={closeModal} />
      )}
    </div>
  );
};

export default TicketList;
