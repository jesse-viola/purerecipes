import recipe_scrapers

from .services import IScrape
from .schemas import Recipe


class ScrapeManager(IScrape):
    def get_recipe(self, link: str) -> Recipe:
        """Uses the recipe-scrapers module to scrape website link into a Recipe schema. :py:class:`~recipe-scrapers._schemaorg.SchemaOrg`

        Args:
            link (str): website url / link

        Returns:
            Recipe: :py:class:`~server.scrape.schemas.Recipe`
        """
        recipe_data = recipe_scrapers.scrape_me(url_path=link, wild_mode=True)
        try: 
            recipe = Recipe(
                title=recipe_data.title(),
                description=recipe_data.description(),
                ingredients=recipe_data.ingredients(),
                instructions=recipe_data.instructions(),
                total_time=recipe_data.total_time(),
                prep_time=recipe_data.prep_time(),
                cook_time=recipe_data.cook_time(),
                image=recipe_data.image(),
                category=recipe_data.category(),
                yields=recipe_data.yields(),
                author=recipe_data.author(),
                language=recipe_data.language(),
                nutrients=recipe_data.nutrients(),
                cuisine=recipe_data.cuisine()
            )
        except recipe_scrapers._exceptions.SchemaOrgException:
            # If one of the fields throws for not found in schema, validation will catch
            pass

        return recipe
