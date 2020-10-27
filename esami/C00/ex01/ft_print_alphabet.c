/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_print_alphabet.c                                :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: abreglia <abreglia@marvin.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/10/23 12:08:15 by abreglia          #+#    #+#             */
/*   Updated: 2020/10/27 00:03:46 by abreglia         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>

void	ft_print_alphabet(void)
{
	char letter;

	letter = 'a';
	while (letter <= 'z')
	{
		write(1, &letter, 1);
		letter++;
	}
}
