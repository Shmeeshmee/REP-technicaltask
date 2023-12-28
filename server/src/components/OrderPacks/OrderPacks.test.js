import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import OrderPacks from './OrderPacks';

describe('<OrderPacks />', () => {
  test('it should mount', () => {
    render(<OrderPacks />);
    
    const orderPacks = screen.getByTestId('OrderPacks');

    expect(orderPacks).toBeInTheDocument();
  });
});