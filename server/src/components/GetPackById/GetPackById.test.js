import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import GetPackById from './GetPackById';

describe('<GetPackById />', () => {
  test('it should mount', () => {
    render(<GetPackById />);
    
    const getPackById = screen.getByTestId('GetPackById');

    expect(getPackById).toBeInTheDocument();
  });
});