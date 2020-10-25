/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_is_negative.c                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: abreglia <abreglia@marvin.fr>              +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2020/10/23 12:45:49 by abreglia          #+#    #+#             */
/*   Updated: 2020/10/23 12:51:17 by abreglia         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <unistd.h>

void	ft_is_negative(int n)
{
	char negative;
	char positive;

	negative = 'N';
	positive = 'P';

	if (n >= 0)
	{
		write(1, &negative, 1);
	} 
	else 
	{
		write(1, &positive, 1);
	}
}
