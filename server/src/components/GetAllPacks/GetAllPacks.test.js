import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import GetAllPacks from './GetAllPacks';

describe('<GetAllPacks />', () => {
  test('it should mount', () => {
    render(<GetAllPacks />);
    
    const getAllPacks = screen.getByTestId('GetAllPacks');

    expect(getAllPacks).toBeInTheDocument();
  });
});