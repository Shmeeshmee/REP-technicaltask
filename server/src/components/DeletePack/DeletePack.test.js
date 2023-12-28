import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import DeletePack from './DeletePack';

describe('<DeletePack />', () => {
  test('it should mount', () => {
    render(<DeletePack />);
    
    const deletePack = screen.getByTestId('DeletePack');

    expect(deletePack).toBeInTheDocument();
  });
});