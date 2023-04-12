from flask import Flask
import dash
import dash_core_components as dcc
import dash_html_components as html
import plotly.express as px
import pandas as pd
from dash import Dash

# Create a Flask instance
server = Flask(__name__)

# Create some sample data
df = pd.DataFrame({'x': [1, 2, 3, 4, 5], 'y': [1, 4, 2, 3, 5]})

# Create a Dash app instance
app = Dash(__name__, server=server, url_base_pathname='/dash/')

# Define the Dash app layout
app.layout = html.Div([
    dcc.Graph(id='scatter-plot', figure={}),
    dcc.Dropdown(
        id='dropdown',
        options=[{'label': str(x), 'value': x} for x in df['x'].unique()],
        value=1
    )
])

# Define a callback function to update the plot based on user input
@app.callback(
    dash.dependencies.Output('scatter-plot', 'figure'),
    [dash.dependencies.Input('dropdown', 'value')]
)
def update_scatter_plot(selected_value):
    # Filter the data based on the selected value
    filtered_df = df[df['x'] == selected_value]
    # Create a Plotly scatter plot
    fig = px.scatter(filtered_df, x='x', y='y')
    # Return the plot
    return fig

# Run the application
if __name__ == '__main__':
    app.run_server(debug=True)
