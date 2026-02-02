import { festivals } from '$lib/data/festivals';
import { getMemoriesByFestival } from '$lib/data/memories';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
	const festival = festivals.find((f) => f.slug === params.slug);

	if (!festival) {
		error(404, {
			message: 'Festival not found'
		});
	}

	// Get memories for this festival
	const memories = getMemoriesByFestival(festival.id);

	// Get related festivals (same heritage type, excluding current)
	const relatedFestivals = festivals
		.filter((f) => f.heritageType === festival.heritageType && f.id !== festival.id)
		.slice(0, 3);

	return {
		festival,
		memories,
		relatedFestivals
	};
};
