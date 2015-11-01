var Root = React.createClass({
  getInitialState: function() {
    return {
      candyNames: [],
      formCandyName: ""
    };
  },
  componentDidMount: function() {
    $.getJSON("/api/candies", function(data) {
      this.setState({candyNames: data["candies"]});
    }.bind(this));
  },
  handleChange: function(evt) {
    evt.preventDefault();
    this.setState({formCandyName: evt.target.value});
  },
  handleSubmit: function(e) {
    e.preventDefault();
    $.ajax({
      url: "/api/candy",
      data: JSON.stringify({"name": this.state.formCandyName}),
      method: "PUT",
      success: function() {
        var newNames = this.state.candyNames;
        newNames.push(this.state.formCandyName);
        this.setState({
          formCandyName: "",
          candyNames: newNames
        });
      }.bind(this)
    });
  },
  render: function() {
    var listElts = [];
    for(var i = 0; i < this.state.candyNames.length; i++) {
      var name = this.state.candyNames[i];
      listElts.push(<li key={name} className="list-group-item">{name}</li>);
    }
    return (
      <div className="container candies">
        <div className="row">
          <ul className="list-group">
            {listElts}
          </ul>
        </div>
        <div className="row">
          <h3>Create New Candy</h3>
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label for="candyName">Candy Name</label>
              <input type="text" className="form-control" value={this.state.formCandyName} id="candyName" placeholder="Candy Name" onChange={this.handleChange}/>
            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
});

ReactDOM.render(<Root/>, document.getElementById("page-content"));
