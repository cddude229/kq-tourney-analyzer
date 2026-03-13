import requests

def fetch_all_pages(
    url,
    headers=None,
    params=None,
    timeout=10,
):
    """
    Fetch all results from a Django REST Framework paginated API.
    
    :param url: Initial API endpoint
    :param headers: Optional request headers (auth tokens, etc.)
    :param params: Optional query params
    :param timeout: Request timeout in seconds
    :return: List of all results
    """
    results = []
    next_url = url

    while next_url:
        response = requests.get(
            "https://kqhivemind.com" + next_url,
            headers=headers,
            params=params,
            timeout=timeout,
        )
        response.raise_for_status()

        data = response.json()

        results.extend(data.get("results", []))
        next_url = data.get("next")

        # After first request, params are already baked into `next`
        params = None

    return results


all_items = fetch_all_pages("/api/tournament/tournament/")

for item in all_items:
    print(item.get("id"), item.get("name"))