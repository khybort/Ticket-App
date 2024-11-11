import React from 'react';

interface UserInputProps {
  userId: string;
  onChange: (value: string) => void;
}

const UserInput: React.FC<UserInputProps> = ({ userId, onChange }) => {
  return (
    <div className="mb-4">
      <label
        htmlFor="user-id"
        className="block text-sm font-semibold text-gray-700 mb-2"
      >
        User ID:
      </label>
      <input
        id="user-id"
        type="text"
        value={userId}
        onChange={(e) => onChange(e.target.value)}
        placeholder="Enter User ID"
        className="w-full px-4 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
    </div>
  );
};

export default UserInput;
