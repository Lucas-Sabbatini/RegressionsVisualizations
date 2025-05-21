import plotly.graph_objects as go

def get_plot_dict():
    fig = go.Figure(data=go.Bar(x=["A", "B", "C"], y=[5, 3, 8]))
    return fig.to_dict()
