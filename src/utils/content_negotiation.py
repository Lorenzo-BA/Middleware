from flask import make_response, request, jsonify
import yaml


def negotiate_content(data, status_code):
    accepted_format = request.accept_mimetypes.best_match(['application/json', 'application/yaml'])

    if accepted_format == 'application/yaml':
        yaml_data = yaml.dump(data, default_flow_style=False)
        response = make_response(yaml_data, status_code)
        response.headers['Content-Type'] = 'application/yaml'
        return response

    return jsonify(data), status_code
